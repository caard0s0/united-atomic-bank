package db

import (
	"context"
	"time"
)

// TransferLoanTransactionParams contains the input parameters of the transfer loan
type TransferLoanTransactionParams struct {
	AccountID    int64     `json:"account_id"`
	LoanAmount   int64     `json:"loan_amount"`
	InterestRate int64     `json:"interest_rate"`
	Status       string    `json:"status"`
	EndDate      time.Time `json:"end_date"`
}

// TransferLoanTransactionResult is the result of the transfer loan
type TransferLoanTransactionResult struct {
	Loan      Loan    `json:"loan"`
	ToAccount Account `json:"to_account"`
	ToEntry   Entry   `json:"to_entry"`
}

// TransferLoanTransaction performs a money transfer to an account.
// It creates a transfer record, add account entry, and update account balance within a single database transaction
func (store *SQLStore) TransferLoanTransaction(ctx context.Context, arg TransferLoanTransactionParams) (TransferLoanTransactionResult, error) {
	var result TransferLoanTransactionResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Loan, err = q.CreateLoan(ctx, CreateLoanParams{
			AccountID:    arg.AccountID,
			LoanAmount:   arg.LoanAmount,
			InterestRate: arg.InterestRate,
			Status:       arg.Status,
			EndDate:      time.Now().UTC().Add(time.Minute),
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.AccountID,
			Amount:    arg.LoanAmount,
		})
		if err != nil {
			return err
		}

		result.ToAccount, err = updateAccountLoanMoney(ctx, q, arg.AccountID, arg.LoanAmount)
		if err != nil {
			return err
		}
		return nil
	})
	return result, err
}

func updateAccountLoanMoney(
	ctx context.Context,
	q *Queries,
	accountID int64,
	amount int64,
) (account Account, err error) {
	account, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID,
		Amount: amount,
	})
	if err != nil {
		return
	}
	return
}
