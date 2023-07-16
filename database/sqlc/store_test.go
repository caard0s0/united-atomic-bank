package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTransferTransactionBetweenAccounts(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before", account1.Balance, account2.Balance)

	// run a concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTransactionBetweenAccountsResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferTransactionBetweenAccounts(ctx, TransferTransactionBetweenAccountsParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	existed := make(map[int]bool)

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)

		// check entries
		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccountID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account2.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check accounts
		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, account1.ID, fromAccount.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account2.ID, toAccount.ID)

		// check accounts' balance
		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)
		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance
		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1%amount == 0)

		k := int(diff1 / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existed, k)
		existed[k] = true
	}

	// check the final updated balances
	updatedAccount1, err := TestQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := TestQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+int64(n)*amount, updatedAccount2.Balance)
}

func TestTransferTransactionBetweenAccountsDeadLock(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	fmt.Println(">> before", account1.Balance, account2.Balance)

	// run a concurrent transfer transactions
	n := 10
	amount := int64(10)

	errs := make(chan error)

	for i := 0; i < n; i++ {
		fromAccountID := account1.ID
		toAccountID := account2.ID

		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}

		go func() {
			ctx := context.Background()
			_, err := store.TransferTransactionBetweenAccounts(ctx, TransferTransactionBetweenAccountsParams{
				FromAccountID: fromAccountID,
				ToAccountID:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balances
	updatedAccount1, err := TestQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := TestQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)
}

func TestTransferLoanToAnAccount(t *testing.T) {
	store := NewStore(testDB)

	account := createRandomAccount(t)
	fmt.Println(">> before", account.Balance)

	// run a concurrent transfer transactions
	n := 5
	amount := int64(10)
	interestRate := int64(1)
	status := "Active"

	errs := make(chan error)
	results := make(chan TransferLoanToAnAccountResult)

	for i := 0; i < n; i++ {
		go func() {
			ctx := context.Background()
			result, err := store.TransferLoanToAnAccount(ctx, TransferLoanToAnAccountParams{
				AccountID:    account.ID,
				LoanAmount:   amount,
				InterestRate: interestRate,
				Status:       status,
				EndDate:      time.Now().Add(time.Minute),
			})

			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer loan
		loan := result.Loan
		require.NotEmpty(t, loan)
		require.Equal(t, account.ID, loan.AccountID)
		require.Equal(t, amount, loan.LoanAmount)
		require.Equal(t, interestRate, loan.InterestRate)
		require.Equal(t, status, loan.Status)
		require.NotZero(t, loan.ID)
		require.NotZero(t, loan.StartDate)
		require.NotZero(t, loan.EndDate)

		_, err = store.GetLoan(context.Background(), loan.ID)
		require.NoError(t, err)

		// check entry
		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account.ID, toEntry.AccountID)
		require.Equal(t, amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// check account
		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, account.ID, toAccount.ID)
	}

	// check the final updated balance
	updatedAccount, err := TestQueries.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)

	fmt.Println(">> after", updatedAccount.Balance)
	require.Equal(t, account.Balance+int64(n)*amount, updatedAccount.Balance)
}
