// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: loan_transfer.sql

package db

import (
	"context"
)

const createLoanTransfer = `-- name: CreateLoanTransfer :one
INSERT INTO loan_transfers (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING id, account_id, amount, interest_rate, open, start_at, end_at
`

type CreateLoanTransferParams struct {
	AccountID int64 `json:"account_id"`
	Amount    int64 `json:"amount"`
}

func (q *Queries) CreateLoanTransfer(ctx context.Context, arg CreateLoanTransferParams) (LoanTransfer, error) {
	row := q.db.QueryRowContext(ctx, createLoanTransfer, arg.AccountID, arg.Amount)
	var i LoanTransfer
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.InterestRate,
		&i.Open,
		&i.StartAt,
		&i.EndAt,
	)
	return i, err
}

const getLoanTransfer = `-- name: GetLoanTransfer :one
SELECT id, account_id, amount, interest_rate, open, start_at, end_at FROM loan_transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetLoanTransfer(ctx context.Context, id int64) (LoanTransfer, error) {
	row := q.db.QueryRowContext(ctx, getLoanTransfer, id)
	var i LoanTransfer
	err := row.Scan(
		&i.ID,
		&i.AccountID,
		&i.Amount,
		&i.InterestRate,
		&i.Open,
		&i.StartAt,
		&i.EndAt,
	)
	return i, err
}