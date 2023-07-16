package db

import (
	"context"
	"testing"
	"time"

	"github.com/caard0s0/united-atomic-bank/util"

	"github.com/stretchr/testify/require"
)

func createRandomLoan(t *testing.T, account Account) Loan {
	arg := CreateLoanParams{
		AccountID:    account.ID,
		LoanAmount:   util.RandomMoney(),
		InterestRate: 1.0,
		Status:       "Active",
		EndDate:      time.Now().Add(time.Minute),
	}

	loan, err := TestQueries.CreateLoan(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, loan)

	require.Equal(t, arg.AccountID, loan.AccountID)
	require.Equal(t, arg.LoanAmount, loan.LoanAmount)
	require.Equal(t, arg.InterestRate, loan.InterestRate)
	require.Equal(t, arg.Status, loan.Status)

	require.NotZero(t, loan.ID)
	require.NotZero(t, loan.StartDate)
	require.NotZero(t, loan.EndDate)

	return loan
}

func TestCreateLoan(t *testing.T) {
	account := createRandomAccount(t)
	createRandomLoan(t, account)
}

func TestGetLoan(t *testing.T) {
	account := createRandomAccount(t)
	loan1 := createRandomLoan(t, account)

	loan2, err := TestQueries.GetLoan(context.Background(), loan1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, loan2)

	require.Equal(t, loan1.ID, loan2.ID)
	require.Equal(t, loan1.AccountID, loan2.AccountID)
	require.Equal(t, loan1.LoanAmount, loan2.LoanAmount)
	require.Equal(t, loan1.InterestRate, loan2.InterestRate)
	require.WithinDuration(t, loan1.StartDate, loan2.StartDate, time.Second)
	require.WithinDuration(t, loan1.EndDate, loan2.EndDate, time.Second)
}
