package database

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
)

var txKey = struct{}{}

type tx struct {
	infra.TxHandler
}

type Dao interface {
	Execute(context.Context, string, ...interface{}) (sql.Result, error)
}

func NewTransaction(h infra.TxHandler) database.SqlTransaction {
	return &tx{h}
}

func (t *tx) Do(ctx context.Context, fn database.ExecFunc) (interface{}, error) {
	err := t.Begin(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, &txKey, t)
	v, err := fn(ctx)
	if err != nil {
		t.Rollback()
		return nil, err
	}

	if err := t.Commit(); err != nil {
		t.Rollback()
		return nil, err
	}
	return v, nil
}

func GetTx(ctx context.Context) (Dao, bool) {
	dao, ok := ctx.Value(&txKey).(infra.TxHandler)
	return dao, ok
}
