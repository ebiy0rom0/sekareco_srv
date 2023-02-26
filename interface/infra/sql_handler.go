package infra

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SqlHandler interface {
	ExecNamedContext(context.Context, string, interface{}) (sql.Result, error)
	GetContext(context.Context, interface{}, string, ...interface{}) error
	QueryRowxContext(context.Context, string, ...interface{}) *sqlx.Row
	SelectContext(context.Context, interface{}, string, ...interface{}) error
	QueryxContext(context.Context, string, ...interface{}) (*sqlx.Rows, error)
}
