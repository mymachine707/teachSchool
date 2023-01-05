package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teachschool/simplebank/util"
)

func createRandomTransfers(t *testing.T, account1, account2 Accaunts) Transfers {

	arg := CreateTransfersParams{
		FromAccauntsID: account1.ID,
		ToAccauntsID:   account2.ID,
		Amount:         util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)
	require.Equal(t, arg.FromAccauntsID, transfer.FromAccauntsID)
	require.Equal(t, arg.ToAccauntsID, transfer.ToAccauntsID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	return transfer
}

func TestCreateTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfers(t, account1, account2)
}

func TestGetTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfers(t, account1, account2)
	transfer2, err := testQueries.GetTransfers(context.Background(), transfer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccauntsID, transfer2.FromAccauntsID)
	require.Equal(t, transfer1.ToAccauntsID, transfer2.ToAccauntsID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)
	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestListTransferss(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomTransfers(t, account1, account2)
	}

	arg := ListTransfersParams{
		FromAccauntsID: account1.ID,
		ToAccauntsID:   account2.ID,
		Limit:          5,
		Offset:         5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, actransfer := range transfers {
		require.NotEmpty(t, actransfer)
	}
}
