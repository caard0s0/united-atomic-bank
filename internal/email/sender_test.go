package email

import (
	"bytes"
	"html/template"
	"testing"
	"time"

	"github.com/caard0s0/united-atomic-bank-server/configs"
	"github.com/caard0s0/united-atomic-bank-server/internal/util"
	"github.com/stretchr/testify/require"
)

func TestSendEmailWithGmail(t *testing.T) {
	config, err := configs.LoadConfig("../..")
	require.NoError(t, err)

	formattedDate := util.FormatDate(time.Now())
	formattedCurrency := util.FormatCurrency("10", "BRL")

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

	tmpl, _ := template.New("emailTemplate").Parse(emailTemplate)

	var body bytes.Buffer
	tmpl.Execute(&body, struct {
		FromAccountOwner string
		ToAccountOwner   string
		Amount           string
		CreatedAt        string
	}{
		FromAccountOwner: "John Doe",
		ToAccountOwner:   "Jane Doe",
		Amount:           formattedCurrency,
		CreatedAt:        formattedDate,
	})

	sender := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)

	subject := "A test email"
	content := body.String()

	to := []string{"vinicardoso216@gmail.com"}
	attachFiles := []string{"../../README.md"}

	err = sender.SendEmail(subject, content, to, nil, nil, attachFiles)
	require.NoError(t, err)
}
