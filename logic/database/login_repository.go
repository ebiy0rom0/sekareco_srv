package database

import "sekareco_srv/domain/model"

type LoginRepository interface {
	Store(model.Login) error
	GetByID(string) (model.Login, error)
}
