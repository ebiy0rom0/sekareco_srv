package sql

import (
	"context"
	"database/sql"
	"sekareco_srv/interface/infra"

	"github.com/ebiy0rom0/errors"
)

// A sqlHandler is database handler wrapper.
//
// [feature]
// Allow switching between different DBMS.
// Only sqlite3 is supported now.
type sqlHandler struct {
	con *sql.DB
}

// NewConnection returns new DB connection.
// First try to connect to MySQL, and if that failure
// switch to a connection to sqlite3.
func NewConnection(user, pass, host, schema string) (*sql.DB, error) {
	db, err := openMysql(user, pass, host, schema)
	if err == nil {
		return db, nil
	}

	db, err = initSqlite3(schema)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return db, nil
}

// NewSqlHandler returns sqlHandler pointer.
func NewSqlHandler(con *sql.DB) *sqlHandler {
	return &sqlHandler{con: con}
}

// Execute returns result at execute argument query.
// Prepared statement are supported, so any argument inject to args.
func (h *sqlHandler) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := h.con.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return res, nil
}

// QueryRow returns 1 record only that result for execute argument query.
// If the query selects no rows, the *sql.Row scan will return ErrNoRows.
func (h *sqlHandler) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// lint:ignore SA5007 too many arguments
	row := h.con.QueryRowContext(ctx, query, args...)
	return row
}

// Query returns rows that result for execute argument query.
func (h *sqlHandler) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// lint:ignore SA5007 too many arguments
	rows, err := h.con.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return rows, nil
}

var _ infra.SqlHandler = (*sqlHandler)(nil)
