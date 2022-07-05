package database

import (
	"context"
	"database/sql"
	"sekareco_srv/logic/database"
)

var txCtxKey = struct{}{}

type tx struct {
	*sql.DB
}

func NewTransaction(db *sql.DB) database.SqlTransaction {
	return &tx{db}
}

func (t *tx) Do(ctx context.Context, fn database.ExecFunc) (interface{}, error) {
	tx, err := t.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		// failed to start transaction
		return nil, err
	}

	ctx = context.WithValue(ctx, &txCtxKey, tx)

	v, err := fn(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}
