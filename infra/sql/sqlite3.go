package sql

import (
	"database/sql"
	"fmt"
	"os"
	"sekareco_srv/util"

	"github.com/ebiy0rom0/errors"
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
		db, err := openSqlite3(source)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		return db, nil
	}

	if err := createDB(source); err != nil {
		return nil, errors.New(err.Error())
	}
	con, err := openSqlite3(source)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if err := createTable(con); err != nil {
		return nil, errors.New(err.Error())
	}

	return con, nil
}

// openSqlite3 establishes a connection with db for sqlite3
// and returns a pointer to the connection.
func openSqlite3(source string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", source)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return db, nil
}

// createDB creates a sqlite3 database file.
// Specify the create location in source.
func createDB(source string) error {
	file, err := os.Create(source)
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()

	return nil
}

// createTable creates a required tables.
// Create a docs/db directory in the root directory and
// stores the queries file for table creation.
func createTable(db *sql.DB) error {
	s := sqlfile.New()

	dir := fmt.Sprintf("%s/%s", util.RootDir(), "docs/db")
	if err := s.Directory(dir); err != nil {
		return errors.New(err.Error())
	}

	if _, err := s.Exec(db); err != nil {
		return errors.New(err.Error())
	}

	return nil
}
