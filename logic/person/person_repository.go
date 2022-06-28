package person

import "sekareco_srv/domain/model"

type PersonRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	RegistPerson(model.Person) (int, error)
	RegistLogin(model.Login) error
	GetPersonByID(int) (model.Person, error)
	GetLoginPerson(string) (model.Login, error)
}
