package sql

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func OpenSqlite3(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDB(dbPath string) error {
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

// for debug
func DropDB(dbPath string) error {
	return os.Remove(dbPath)
}
