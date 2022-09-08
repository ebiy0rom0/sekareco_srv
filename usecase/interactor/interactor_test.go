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
	lr := database.NewLoginRepository(test.InjectSqlHandler())
	ai = interactor.NewAuthInteractor(test.InjectTokenManager(), lr, tx)

	mr := database.NewMusicRepository(test.InjectSqlHandler())
	mi = interactor.NewMusicInteractor(mr, tx)

	m.Run()
}
