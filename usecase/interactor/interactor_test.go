package interactor_test

import (
	"sekareco_srv/interface/database"
	"sekareco_srv/test"
	"sekareco_srv/usecase/interactor"
	"testing"
)

func TestMain(m *testing.M) {
	test.Setup()

	tx := database.NewTransaction(test.InjectTxHandler())
	l := database.NewLoginRepository(test.InjectSqlHandler())
	i = interactor.NewAuthInteractor(test.InjectTokenManager(), l, tx)

	m.Run()
}
