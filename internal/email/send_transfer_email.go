package email

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"time"

	"github.com/caard0s0/united-atomic-bank-server/configs"
	db "github.com/caard0s0/united-atomic-bank-server/internal/database/sqlc"
	"github.com/caard0s0/united-atomic-bank-server/internal/util"
)

func SendEmailWithSuccessfulTransfer(transferResult db.TransferTransactionResult, email string) error {
	config, err := configs.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot read config:", err)
	}

	formattedDate := util.FormatDate(time.Now())
	formattedCurrency := util.FormatCurrency(transferResult.Transfer.Amount, transferResult.FromAccount.Currency)

	emailTemplate := `
		<!DOCTYPE html>
		<html>
		</head>
			<body style="background-color: #E5E5E5; margin: 0; padding: 0;">
				<div style="max-width: 600px; margin: 0 auto; padding-bottom: 2rem; background-color: #fff;">
					<div style="background: linear-gradient(215deg, #171d26 15%, #000 85%);">
						<img style="width: 8rem; height: 8rem;" src="https://github.com/caard0s0/united-atomic-bank-server/assets/95318788/d2d8a5e9-8ba3-48e6-95d6-30d31bb0618e" alt="United Atomic Bank Logo" title="United Atomic Bank Logo">
						<h1 style="color: #fff; padding-bottom: 2rem; margin-left: 2rem; margin-right: 2rem; text-align: left; font-size: 2rem;">Transfer Made</h1>
					</div>
					<div>
						<p style="margin-left: 2rem; margin-right: 2rem; text-align: left; font-size: 1.8rem; margin-top: 4rem; color: #171d26; font-weight: 500;">Hi, {{.FromAccountOwner}}</p>
						<p style="margin-left: 2rem; margin-right: 2rem; text-align: left; font-size: 1rem; color: #555555; margin-bottom: 3rem;">The transfer to <strong style="text-transform: uppercase;">{{.ToAccountOwner}}</strong> was carried out successfully.</p>
						
						<div style="padding: 20px 25%;">
							<div style="font-size: small; text-align: center; padding-top: 2rem; padding-bottom: 2rem; background-color: #E5E5E5; color: #171d26;">
								<p>Amount Sent</p>
								<p>{{.Amount}}</p>
								<p>{{.CreatedAt}}</p>
							</div>
						</div>
					</div>
					<p style="margin-left: 2rem; margin-right: 2rem; text-align: left; margin-top: 3rem; color: #191919; font-size: x-large; font-weight: bold;">Hugs, <br> UAB Team</p>
				</div>
			</body>
		</html>
	`

	tmpl, err := template.New("emailTemplate").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("failed to create new template: %w", err)
	}

	var body bytes.Buffer
	err = tmpl.Execute(&body, struct {
		FromAccountOwner string
		ToAccountOwner   string
		Amount           string
		CreatedAt        string
	}{
		FromAccountOwner: transferResult.Transfer.FromAccountOwner,
		ToAccountOwner:   transferResult.Transfer.ToAccountOwner,
		Amount:           formattedCurrency,
		CreatedAt:        formattedDate,
	})
	if err != nil {
		return fmt.Errorf("failed to execute template: %w", err)
	}

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "Transfer completed successfully"
	content := body.String()

	to := []string{email}

	err = sender.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	return nil
}
