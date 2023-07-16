package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	Querier
	TransferTransactionBetweenAccounts(ctx context.Context, arg TransferTransactionBetweenAccountsParams) (TransferTransactionBetweenAccountsResult, error)
	TransferLoanToAnAccount(ctx context.Context, arg TransferLoanToAnAccountParams) (TransferLoanToAnAccountResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	db *sql.DB
	*Queries
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTransactionBetweenAccountsParams contains the input parameters of the transfer transaction
type TransferTransactionBetweenAccountsParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTransactionBetweenAccountsResult is the result of the transfer transaction
type TransferTransactionBetweenAccountsResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferLoanToAnAccountParams contains the input parameters of the transfer loan
type TransferLoanToAnAccountParams struct {
	AccountID    int64     `json:"account_id"`
	LoanAmount   int64     `json:"loan_amount"`
	InterestRate int64     `json:"interest_rate"`
	Status       string    `json:"status"`
	EndDate      time.Time `json:"end_date"`
}

// TransferLoanToAnAccountResult is the result of the transfer loan
type TransferLoanToAnAccountResult struct {
	Loan      Loan    `json:"loan"`
	ToAccount Account `json:"to_account"`
	ToEntry   Entry   `json:"to_entry"`
}

// TransferTransactionBetweenAccounts performs a money transfer from one account to the other.
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *SQLStore) TransferTransactionBetweenAccounts(ctx context.Context, arg TransferTransactionBetweenAccountsParams) (TransferTransactionBetweenAccountsResult, error) {
	var result TransferTransactionBetweenAccountsResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccount, result.ToAccount, err = updateAccountMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
			if err != nil {
				return err
			}
		} else {
			result.ToAccount, result.FromAccount, err = updateAccountMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return result, err
}

// TransferLoanToAnAccount performs a money transfer to an account.
// It creates a transfer record, add account entry, and update account balance within a single database transaction
func (store *SQLStore) TransferLoanToAnAccount(ctx context.Context, arg TransferLoanToAnAccountParams) (TransferLoanToAnAccountResult, error) {
	var result TransferLoanToAnAccountResult

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

func updateAccountMoney(
	ctx context.Context,
	q *Queries,
	accountID1 int64,
	amount1 int64,
	accountID2 int64,
	amount2 int64,
) (account1 Account, account2 Account, err error) {
	account1, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID1,
		Amount: amount1,
	})
	if err != nil {
		return
	}

	account2, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		ID:     accountID2,
		Amount: amount2,
	})
	return
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
