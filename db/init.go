package db

import (
	"database/sql"
	"os"

	"github.com/tanimutomo/sqlfile"
)

func Init() error {
	_, err := os.Stat(DATABASE_SCHEMA_NAME)
	if err != nil {
		err = CreateDB()
		if err != nil {
			return err
		}

		db, err := OpenSqlite3()
		if err != nil {
			return err
		}
		defer db.Close()

		err = CreateTable(db)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateTable(db *sql.DB) error {
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
