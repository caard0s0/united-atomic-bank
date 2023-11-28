package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg CreateTransferParams) (TransferTxResult, error)
	LoanTransferTransaction(ctx context.Context, arg CreateLoanTransferParams) (LoanTransferTransactionResult, error)
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
