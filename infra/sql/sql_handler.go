package sql

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"

	"github.com/ebiy0rom0/errors"
	"github.com/jmoiron/sqlx"
)

// A sqlHandler is database handler wrapper.
//
// [feature]
// Allow switching between different DBMS.
// Only sqlite3 is supported now.
type sqlHandler struct {
	con *sqlx.DB
}

// NewConnection returns new DB connection.
// First try to connect to MySQL, and if that failure
// switch to a connection to sqlite3.
func NewConnection(user, pass, host, schema string) (*sqlx.DB, error) {
	db, err := openMysql(user, pass, host, schema)
	if err == nil {
		return db, nil
	}

	db, err = initSqlite3(schema)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}

// NewSqlHandler returns sqlHandler pointer.
func NewSqlHandler(con *sqlx.DB) *sqlHandler {
	return &sqlHandler{con: con}
}

// Execute returns result at execute argument query.
// Any named placeholder parameters are replaced with fields from arg.
func (h *sqlHandler) ExecNamedContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	stmt, err := h.con.PrepareNamedContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, arg)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return res, nil
}

// GetContext does a QueryRow and scans the result row to dest.
// If the query selects no row, scan will return ErrNoRows.
func (h *sqlHandler) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if err := h.con.GetContext(ctx, dest, query, args...); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// QueryRowxContext does a QueryRow and returns *sqlx.Row.
// If the query selects no row, scan will return ErrNoRows.
func (h *sqlHandler) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return h.con.QueryRowxContext(ctx, query, args...)
}

// SelectContext does a Query and scans each row into dest.
func (h *sqlHandler) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	if err := h.con.SelectContext(ctx, dest, query, args...); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// SelectContext does a Query and returns *sqlx.Rows.
func (h *sqlHandler) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	rows, err := h.con.QueryxContext(ctx, query, args...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return rows, nil
}

var _ infra.SqlHandler = (*sqlHandler)(nil)
