package db

import (
	"context"
	"testing"
	"time"

	"github.com/caard0s0/united-atomic-bank-server/util"

	"github.com/stretchr/testify/require"
)

func createRandomLoanTransfer(t *testing.T, account Account) LoanTransfer {
	arg := CreateLoanTransferParams{
		AccountID: account.ID,
		Amount:    util.RandomMoney(),
	}

	loan, err := TestQueries.CreateLoanTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, loan)

	require.Equal(t, arg.AccountID, loan.AccountID)
	require.Equal(t, arg.Amount, loan.Amount)

	require.NotZero(t, loan.ID)
	require.NotZero(t, loan.StartAt)
	require.NotZero(t, loan.EndAt)

	return loan
}

func TestCreateLoanTransfer(t *testing.T) {
	account := createRandomAccount(t)
	createRandomLoanTransfer(t, account)
}

func TestGetLoanTransfer(t *testing.T) {
	account := createRandomAccount(t)
	loan1 := createRandomLoanTransfer(t, account)

	loan2, err := TestQueries.GetLoanTransfer(context.Background(), loan1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, loan2)

	require.Equal(t, loan1.ID, loan2.ID)
	require.Equal(t, loan1.AccountID, loan2.AccountID)
	require.Equal(t, loan1.Amount, loan2.Amount)
	require.Equal(t, loan1.InterestRate, loan2.InterestRate)
	require.WithinDuration(t, loan1.StartAt, loan2.StartAt, time.Second)
	require.WithinDuration(t, loan1.EndAt, loan2.EndAt, time.Second)
}
