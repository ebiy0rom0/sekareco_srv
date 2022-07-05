package infra

import (
	"database/sql"
	"os"

	_sql "sekareco_srv/infra/sql"

	"github.com/tanimutomo/sqlfile"
)

type SqlHandler struct {
	Conn *sql.DB
}

func NewSqlHandler(dbPath string) (h *SqlHandler, err error) {
	var db *sql.DB

	if _, err = os.Stat(dbPath); err != nil {
		if err = _sql.CreateDB(dbPath); err != nil {
			return
		}

		if db, err = _sql.OpenSqlite3(dbPath); err != nil {
			return
		}

		if err = createTable(db); err != nil {
			return
		}

	} else {
		if db, err = _sql.OpenSqlite3(dbPath); err != nil {
			return
		}
	}

	h = &SqlHandler{
		Conn: db,
	}
	return
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
