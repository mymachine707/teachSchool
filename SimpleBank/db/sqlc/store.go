package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
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

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfers `json:"transfer"`
	FromAccaunt Accaunts  `json:"from_account"`
	ToAccaunt   Accaunts  `json:"to_account"`
	FromEntry   Enteries  `json:"from_entry"`
	ToEntry     Enteries  `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfers(ctx, CreateTransfersParams{
			FromAccauntsID: arg.FromAccountID,
			ToAccauntsID:   arg.ToAccountID,
			Amount:         arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEnteries(ctx, CreateEnteriesParams{
			AccauntsID: arg.FromAccountID,
			Amount:     -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEnteries(ctx, CreateEnteriesParams{
			AccauntsID: arg.ToAccountID,
			Amount:     arg.Amount,
		})
		if err != nil {
			return err
		}
		if arg.FromAccountID < arg.ToAccountID {
			result.FromAccaunt, result.ToAccaunt, err = addMoney(ctx, q, arg.FromAccountID, -arg.Amount, arg.ToAccountID, arg.Amount)
		} else {
			result.ToAccaunt, result.FromAccaunt, err = addMoney(ctx, q, arg.ToAccountID, arg.Amount, arg.FromAccountID, -arg.Amount)
		}
		// Todo update account balanse

		return nil

	})

	return result, err
}

func addMoney(
	ctx context.Context,
	q *Queries,
	acccountID1 int64,
	amount1 int64,
	acccountID2 int64,
	amount2 int64,
) (account1 Accaunts, account2 Accaunts, err error) {
	account1, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     acccountID1,
		Amount: amount1,
	})

	if err != nil {
		return
	}
	account2, err = q.AddAccountBalance(ctx, AddAccountBalanceParams{
		ID:     acccountID2,
		Amount: amount2,
	})

	if err != nil {
		return
	}

	return
}
