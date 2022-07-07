package database

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
)

var txCtxKey = struct{}{}

type tx struct {
	infra.TxHandler
}

func NewTransaction(h infra.TxHandler) database.SqlTransaction {
	return &tx{h}
}

func (t *tx) Do(ctx context.Context, fn database.ExecFunc) (interface{}, error) {
	ctx = context.WithValue(ctx, &txCtxKey, t)

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
