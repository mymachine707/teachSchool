package db

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teachschool/simplebank/util"
)

func createRandomEnteries(t *testing.T, account Accaunts) Enteries {
	arg := CreateEnteriesParams{
		AccauntsID: account.ID,
		Amount:     util.RandomMoney(),
	}

	enteries, err := testQueries.CreateEnteries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, enteries)
	require.Equal(t, arg.AccauntsID, enteries.AccauntsID)
	require.Equal(t, arg.Amount, enteries.Amount)

	require.NotZero(t, enteries.ID)
	require.NotZero(t, enteries.CreatedAt)
	return enteries
}

func TestCreateEnteries(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEnteries(t, account)
}

func TestGetEnteries(t *testing.T) {
	account1 := createRandomAccount(t)
	entry1 := createRandomEnteries(t, account1)
	entry2, err := testQueries.GetEnteries(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccauntsID, entry2.AccauntsID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEnteriess(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEnteries(t, account)
	}

	arg := ListEnteriessParams{
		AccauntsID: account.ID,
		Limit:      5,
		Offset:     5,
	}

	enteriess, err := testQueries.ListEnteriess(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, enteriess, 5)

	for _, acenteries := range enteriess {
		require.NotEmpty(t, acenteries)
	}
}
