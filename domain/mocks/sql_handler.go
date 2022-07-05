package mocks

import (
	"sekareco_srv/infra"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewMockSqlHandler() infra.SqlHandler {
	db, mock, err := sqlmock.New()
	return infra.SqlHandler{
		Conn: db,
		Tx:   nil,
	}
}
