package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})

			errs <- err       //?!!!
			results <- result //?!!!
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		results := <-results
		require.NotEmpty(t, results)

		transfer := results.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccauntsID)
		require.Equal(t, account2.ID, transfer.ToAccauntsID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		_, err = store.GetTransfers(context.Background(), transfer.ID)
		require.NoError(t, err)

		fromEntry := results.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, account1.ID, fromEntry.AccauntsID)
		require.Equal(t, -amount, fromEntry.Amount)
		require.NotZero(t, fromEntry.ID)
		require.NotZero(t, fromEntry.CreatedAt)

		_, err = store.GetEnteries(context.Background(), fromEntry.ID)
		require.NoError(t, err)

		toEntry := results.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, account1.ID, toEntry.AccauntsID)
		require.Equal(t, -amount, toEntry.Amount)
		require.NotZero(t, toEntry.ID)
		require.NotZero(t, toEntry.CreatedAt)

		_, err = store.GetEnteries(context.Background(), toEntry.ID)
		require.NoError(t, err)

		// todo check account balance
	}
}
