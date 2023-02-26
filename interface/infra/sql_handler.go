package infra

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SqlHandler interface {
	Execute(context.Context, string, ...interface{}) (sql.Result, error)
	QueryRow(context.Context, string, ...interface{}) *sqlx.Row
	Query(context.Context, string, ...interface{}) (*sqlx.Rows, error)
}
