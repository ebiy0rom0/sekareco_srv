package infra

import (
	"database/sql"
	"os"

	"github.com/tanimutomo/sqlfile"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler() (*SqlHandler, error) {
	var db *sql.DB

	_, err := os.Stat(DATABASE_SCHEMA_NAME)
	if err != nil {
		err = CreateDB()
		if err != nil {
			return nil, err
		}

		db, err = OpenSqlite3()
		if err != nil {
			return nil, err
		}

		err = createTable(db)
		if err != nil {
			return nil, err
		}

	} else {
		db, err = OpenSqlite3()
		if err != nil {
			return nil, err
		}
	}

	handler := new(SqlHandler)
	handler.Conn = db

	return handler, nil
}

func createTable(db *sql.DB) error {
	s := sqlfile.New()

	err := s.Directory("./doc/db")
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
