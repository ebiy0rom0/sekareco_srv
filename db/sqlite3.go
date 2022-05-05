package db

import (
	"database/sql"
	"os"
)

const DATABASE_SCHEMA_NAME = "./sekareco.db"

func InitSqlite3() error {
	return nil
}

func OpenSqlite3() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", DATABASE_SCHEMA_NAME)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func CreateDB() error {
	file, err := os.Create(DATABASE_SCHEMA_NAME)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}
