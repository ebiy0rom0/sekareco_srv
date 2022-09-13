package sql

import (
	"database/sql"
	"fmt"
	"os"
	"sekareco_srv/util"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tanimutomo/sqlfile"
)

// initSqlite3 returns a pointer to the connection.
// If first time and not exist data source that
// create database file and create the need tables.
func initSqlite3(schema string) (*sql.DB, error) {
	var con *sql.DB

	// You need to making db/ in the root directory.
	source := fmt.Sprintf("%s/%s/%s", util.RootDir(), os.Getenv("DB_PATH"), schema)

	if _, err := os.Stat(source); err == nil {
		return openSqlite3(source)
	}

	if err := createDB(source); err != nil {
		return nil, err
	}
	con, err := openSqlite3(source)
	if err != nil {
		return nil, err
	}
	if err := createTable(con); err != nil {
		return nil, err
	}

	return con, nil
}

// openSqlite3 establishes a connection with db for sqlite3
// and returns a pointer to the connection.
func openSqlite3(source string) (*sql.DB, error) {
	return sql.Open("sqlite3", source)
}

func createDB(source string) error {
	file, err := os.Create(source)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

// createTable
func createTable(db *sql.DB) error {
	s := sqlfile.New()

	dir := fmt.Sprintf("%s/%s", util.RootDir(), "docs/db")
	if err := s.Directory(dir); err != nil {
		return err
	}

	if _, err := s.Exec(db); err != nil {
		return err
	}

	return nil
}
