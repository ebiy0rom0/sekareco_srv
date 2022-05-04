package db

import (
	"database/sql"
	"fmt"
)

func OpenMySQL() (*sql.DB, error) {
	dataSorceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		"account", "password", "host:port", "db_schema",
	)

	db, err := sql.Open("mysql", dataSorceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
