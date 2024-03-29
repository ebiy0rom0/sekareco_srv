package database_test

import (
	"sekareco_srv/interface/database"
	"sekareco_srv/test"
	"testing"
)

func TestMain(m *testing.M) {
	test.Initialize()

	loginRepo = database.NewLoginRepository(test.InjectSqlHandler())
	musicRepo = database.NewMusicRepository(test.InjectSqlHandler())
	personRepo = database.NewPersonRepository(test.InjectSqlHandler())
	recordRepo = database.NewRecordRepository(test.InjectSqlHandler())

	tr = database.NewTransaction(test.InjectTxHandler())

	m.Run()
}
