package infra

import (
	"context"
	"database/sql"
)

type TxHandler interface {
	Begin(context.Context, *sql.TxOptions) error
	Execute(context.Context, string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}
