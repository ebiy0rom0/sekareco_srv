package infra

import "database/sql"

type TxHandler interface {
	Execute(string, ...interface{}) (sql.Result, error)
	Commit() error
	Rollback() error
}
