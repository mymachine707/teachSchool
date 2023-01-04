package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/teachschool/simplebank/util"
)

// accaunts_id,
// amount

func createRandomEnteries(t *testing.T) Enteries {
	arg := CreateEnteriesParams{
		AccauntsID: util.RandomOwner(),
		Amount:     util.RandomMoney(),
	}

	enteries, err := testQueries.CreateEnteries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, enteries)
	require.Equal(t, arg.Owner, enteries.Owner)
	require.Equal(t, arg.Balance, enteries.Balance)
	require.Equal(t, arg.Currency, enteries.Currency)

	require.NotZero(t, enteries.ID)
	require.NotZero(t, enteries.CreatedAt)
	return enteries
}

func TestCreateEnteries(t *testing.T) {
	createRandomEnteries(t)
}

func TestGetEnteries(t *testing.T) {
	enteries1 := createRandomEnteries(t)
	enteries2, err := testQueries.GetEnteries(context.Background(), enteries1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, enteries2)

	require.Equal(t, enteries1.ID, enteries2.ID)
	require.Equal(t, enteries1.Owner, enteries2.Owner)
	require.Equal(t, enteries1.Balance, enteries2.Balance)
	require.Equal(t, enteries1.Currency, enteries2.Currency)
	require.WithinDuration(t, enteries1.CreatedAt, enteries2.CreatedAt, time.Second)
}

func TestUpdateEnteries(t *testing.T) {
	enteries1 := createRandomEnteries(t)

	arg := UpdateEnteriesParams{
		ID:      enteries1.ID,
		Balance: util.RandomMoney(),
	}
	enteries2, err := testQueries.UpdateEnteries(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, enteries2)

	require.Equal(t, enteries1.ID, enteries2.ID)
	require.Equal(t, enteries1.Owner, enteries2.Owner)
	require.Equal(t, arg.Balance, enteries2.Balance)
	require.Equal(t, enteries1.Currency, enteries2.Currency)
	require.WithinDuration(t, enteries1.CreatedAt, enteries2.CreatedAt, time.Second)
}

func TestDeleteEnteries(t *testing.T) {
	enteries1 := createRandomEnteries(t)
	err := testQueries.DeleteEnteries(context.Background(), enteries1.ID)
	require.NoError(t, err)

	enteries2, err := testQueries.GetEnteries(context.Background(), enteries1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, enteries2)
}

func TestListEnteriess(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEnteries(t)
	}

	arg := ListEnteriessParams{
		Limit:  5,
		Offset: 5,
	}

	enteriess, err := testQueries.ListEnteriess(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, enteriess, 5)

	for _, acenteries := range enteriess {
		require.NotEmpty(t, acenteries)
	}
}
