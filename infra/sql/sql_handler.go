package sql

import (
	"context"
	"database/sql"
	"os"

	"sekareco_srv/interface/infra"
)

var txCtxKey = struct{}{}

type SqlHandler struct {
	Con *sql.DB
}

func NewSqlHandler(dbPath string) (h *SqlHandler, err error) {
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

	h = &SqlHandler{db}
	return
}

func (h *SqlHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	stmt, err := h.Con.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(args...)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (h *SqlHandler) QueryRow(query string, args ...interface{}) *sql.Row {
	// lint:ignore SA5007 too many argments
	row := h.Con.QueryRow(query, args...)
	return row
}

func (h *SqlHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	// lint:ignore SA5007 too many argments
	rows, err := h.Con.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (h *SqlHandler) BeginTx(ctx context.Context) (infra.TxHandler, error) {
	tx, err := h.Con.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		// failed to start transaction
		return nil, err
	}

	return newTxHandler(tx), nil
}
