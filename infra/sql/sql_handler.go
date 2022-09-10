package sql

import (
	"context"
	"database/sql"
	"os"
)

// A sqlHandler is database handler wrapper.
// [feature]
// Allow switching between different DBMS.
// Only sqlite3 is supported now.
type sqlHandler struct {
	con *sql.DB
}

// NewSqlHandler returns sqlHandler and txHandler pointer.
// If not exists sqliteDB, create database and migrate require tables
// before connection opener.
func NewSqlHandler(dbPath string) (h *sqlHandler, th *txHandler, err error) {
	var db *sql.DB

	if _, err = os.Stat(dbPath); err != nil {
		if err = createDB(dbPath); err != nil {
			return
		}

		if db, err = openSqlite3(dbPath); err != nil {
			return
		}

		if err = createTable(db); err != nil {
			return
		}

	} else {
		if db, err = openSqlite3(dbPath); err != nil {
			return
		}
	}

	h = &sqlHandler{con: db}
	th = &txHandler{con: db}
	return
}

func (h *sqlHandler) Execute(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	stmt, err := h.con.PrepareContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (h *sqlHandler) QueryRow(ctx context.Context, query string, args ...interface{}) *sql.Row {
	// lint:ignore SA5007 too many arguments
	row := h.con.QueryRowContext(ctx, query, args...)
	return row
}

func (h *sqlHandler) Query(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	// lint:ignore SA5007 too many arguments
	rows, err := h.con.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
