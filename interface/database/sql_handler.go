package database

import "database/sql"

type SqlHandler interface {
	Execute(string, ...interface{}) (sql.Result, error)
	QueryRow(string, ...interface{}) *sql.Row
	Query(string, ...interface{}) (*sql.Rows, error)
	StartTransaction() error
	Commit() error
	Rollback() error
}
