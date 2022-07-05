package database

import "sekareco_srv/domain/model"

type PersonRepository interface {
	Store(model.Person) (int, error)
	GetByID(int) (model.Person, error)
}
