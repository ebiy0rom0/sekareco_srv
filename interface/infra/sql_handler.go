package infra

import (
	"context"
	"database/sql"
)

type SqlHandler interface {
	Execute(context.Context, string, ...interface{}) (sql.Result, error)
	QueryRow(context.Context, string, ...interface{}) *sql.Row
	Query(context.Context, string, ...interface{}) (*sql.Rows, error)
}
