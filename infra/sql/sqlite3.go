package sql

import (
	"database/sql"
	"os"
	"sekareco_srv/util"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tanimutomo/sqlfile"
)

func openSqlite3(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createDB(dbPath string) error {
	file, err := os.Create(dbPath)
	if err != nil {
		return err
	}
	file.Close()

	return nil
}

func createTable(db *sql.DB) error {
	s := sqlfile.New()

	err := s.Directory(util.RootDir() + "/docs/db")
	if err != nil {
		return err
	}

	_, err = s.Exec(db)
	if err != nil {
		return err
	}

	return nil
}

// for debug
func DropDB(dbPath string) error {
	return os.Remove(dbPath)
}
