package inputport

import "sekareco_srv/domain/model"

type PersonLogic interface {
	Store(model.PostPerson) (model.Person, error)
	GetByID(int) (model.Person, error)
	IsDuplicateLoginID(string) (bool, error)
}
