package db

import (
	"context"
	"time"
)

// LoanTransferTransactionResult is the result of the transfer loan
type LoanTransferTransactionResult struct {
	Loan      LoanTransfer `json:"loan"`
	ToAccount Account      `json:"to_account"`
	ToEntry   Entry        `json:"to_entry"`
}

// LoanTransferTransaction performs a money transfer to an account.
// It creates a transfer record, add account entry, and update account balance within a single database transaction
func (store *SQLStore) LoanTransferTransaction(ctx context.Context, arg CreateLoanTransferParams) (LoanTransferTransactionResult, error) {
	var result LoanTransferTransactionResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Loan, err = q.CreateLoanTransfer(ctx, CreateLoanTransferParams{
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
