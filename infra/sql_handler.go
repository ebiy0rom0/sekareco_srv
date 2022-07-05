package infra

import (
	"database/sql"
	"fmt"
	"os"

	_sql "sekareco_srv/infra/sql"

	"github.com/tanimutomo/sqlfile"
)

type SqlHandler struct {
	Conn *sql.DB
	Tx   *sql.Tx
}

func NewSqlHandler(dbPath string) (*SqlHandler, error) {
	var db *sql.DB

	_, err := os.Stat(dbPath)
	if err != nil {
		err = _sql.CreateDB(dbPath)
		if err != nil {
			return nil, err
		}

		db, err = _sql.OpenSqlite3(dbPath)
		if err != nil {
			return nil, err
		}

		err = createTable(db)
		if err != nil {
			return nil, err
		}

	} else {
		db, err = _sql.OpenSqlite3(dbPath)
		if err != nil {
			return nil, err
		}
	}

	handler := new(SqlHandler)
	handler.Conn = db
	handler.Tx = nil

	return handler, nil
}

func createTable(db *sql.DB) error {
	s := sqlfile.New()

	err := s.Directory("./../doc/db")
	if err != nil {
		return err
	}

	_, err = s.Exec(db)
	if err != nil {
		return err
	}

	return nil
}

func (h *SqlHandler) Execute(query string, args ...interface{}) (sql.Result, error) {
	if h.Tx == nil {
		return nil, fmt.Errorf("no transaction")
	}

	stmt, err := h.Conn.Prepare(query)
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
	row := h.Conn.QueryRow(query, args...)
	return row
}

func (h *SqlHandler) Query(query string, args ...interface{}) (*sql.Rows, error) {
	rows, err := h.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (h *SqlHandler) Prepare(query string) (*sql.Stmt, error) {
	if h.Tx == nil {
		return nil, fmt.Errorf("no transaction")
	}
	return h.Tx.Prepare(query)
}

func (h *SqlHandler) StartTransaction() (err error) {
	tx, err := h.Conn.Begin()
	if err != nil {
		return
	}

	h.Tx = tx
	return
}

func (h *SqlHandler) Commit() (err error) {
	if h.Tx == nil {
		err = fmt.Errorf("no transaction")
		return
	}

	err = h.Tx.Commit()
	h.Tx = nil
	return
}

func (h *SqlHandler) Rollback() (err error) {
	if h.Tx == nil {
		err = fmt.Errorf("no transaction")
		return
	}

	err = h.Tx.Rollback()
	h.Tx = nil
	return
}