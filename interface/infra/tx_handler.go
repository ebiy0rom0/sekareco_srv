package infra

import (
	"context"
	"database/sql"
)

type Executor interface {
	ExecNamedContext(context.Context, string, interface{}) (sql.Result, error)
	UpdateNamedContext(context.Context, string, interface{}, ...interface{}) (sql.Result, error)
}

type TxHandler interface {
	Executor
	BeginTxx(context.Context, *sql.TxOptions) error
	Commit() error
	Rollback() error
}
