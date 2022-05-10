package logic

import "sekareco_srv/domain"

type PersonRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	RegistPerson(domain.Person) (int, error)
	RegistLogin(domain.Login) error
	GetPersonById(int) (domain.Person, error)
	GetLoginPerson(string) (domain.Login, error)
}
