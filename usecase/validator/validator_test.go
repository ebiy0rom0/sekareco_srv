package validator_test

import (
	"sekareco_srv/interface/database"
	"sekareco_srv/test"
	"sekareco_srv/usecase/validator"
	"testing"
)

func TestMain(m *testing.M) {
	test.Initialize()

	lr := database.NewLoginRepository(test.InjectSqlHandler())
	pv = validator.NewPersonValidator(lr)

	av = validator.NewAuthValidator()

	m.Run()
}
