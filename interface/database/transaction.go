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

type Dao interface {
	Execute(context.Context, string, ...interface{}) (sql.Result, error)
}

func NewTransaction(h infra.TxHandler) *tx {
	return &tx{h}
}

// Do executes the function passed in ExecFunc in a transaction.
// If you want to execute multiple functions in the same transaction,
// wrapped functions in a closure and pass to ExecFunc.
func (t *tx) Do(ctx context.Context, fn database.ExecFunc) (interface{}, error) {
	opt := &sql.TxOptions{Isolation: sql.LevelReadCommitted}
	if err := t.Begin(ctx, opt); err != nil {
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
func getTx(ctx context.Context) (Dao, bool) {
	dao, ok := ctx.Value(&txKey).(infra.TxHandler)
	return dao, ok
}

// interface implementation checks
var _ database.SqlTransaction = (*tx)(nil)
