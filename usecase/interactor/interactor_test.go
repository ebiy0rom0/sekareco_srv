package interactor_test

import (
	_ "sekareco_srv/infra"
	"sekareco_srv/interface/database"
	"sekareco_srv/test"
	"sekareco_srv/usecase/interactor"
	"testing"
)

func TestMain(m *testing.M) {
	test.Initialize()

	tx := database.NewTransaction(test.InjectTxHandler())

	lr := database.NewLoginRepository(test.InjectSqlHandler())
	mr := database.NewMusicRepository(test.InjectSqlHandler())
	pr := database.NewPersonRepository(test.InjectSqlHandler())
	rr := database.NewRecordRepository(test.InjectSqlHandler())

	ai = interactor.NewAuthInteractor(test.InjectTokenManager(), lr, tx)
	mi = interactor.NewMusicInteractor(mr, tx)
	pi = interactor.NewPersonInteractor(pr, lr, tx)

	ri = interactor.NewRecordInteractor(rr, tx)

	m.Run()
}
