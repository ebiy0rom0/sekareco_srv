package database

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"

	"github.com/ebiy0rom0/errors"
)

var txKey = struct{}{}

type tx struct {
	infra.TxHandler
}

func NewTransaction(h infra.TxHandler) database.SqlTransaction {
	return &tx{h}
}

// Do executes the function passed in ExecFunc in a transaction.
// If you want to execute multiple functions in the same transaction,
// wrapped functions in a closure and pass to ExecFunc.
func (t *tx) Do(ctx context.Context, fn database.ExecFunc) (interface{}, error) {
	opt := &sql.TxOptions{Isolation: sql.LevelReadCommitted}
	if err := t.BeginTxx(ctx, opt); err != nil {
		return nil, errors.WithStack(err)
	}

	ctx = context.WithValue(ctx, &txKey, t)
	v, err := fn(ctx)
	if err != nil {
		t.Rollback()
		return nil, errors.WithStack(err)
	}

	if err := t.Commit(); err != nil {
		t.Rollback()
		return nil, errors.WithStack(err)
	}
	return v, nil
}

// getTx returns transaction handler object.
// It can only be retrieved in the function passed by ExecFunc.
func getTx(ctx context.Context) (infra.Executor, bool) {
	dao, ok := ctx.Value(&txKey).(infra.TxHandler)
	return dao, ok
}

// interface implementation checks
var _ database.SqlTransaction = (*tx)(nil)
