package infra

import (
	"context"
	"database/sql"
)

type TxHandler interface {
	BeginTxx(context.Context, *sql.TxOptions) error
	ExecNamedContext(context.Context, string, interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}
