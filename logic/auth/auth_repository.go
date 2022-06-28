package auth

import "sekareco_srv/domain/model"

type AuthRepository interface {
	GetLoginPerson(string) (model.Login, error)
}
