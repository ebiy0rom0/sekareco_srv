package inputport

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
)

type PersonInputport interface {
	Store(context.Context, inputdata.AddPerson) (model.Person, error)
	Update(context.Context, int, inputdata.UpdatePerson) error
	GetByID(context.Context, int) (model.Person, error)
}

type PersonValidator interface {
	ValidationAdd(context.Context, inputdata.AddPerson) error
	ValidationUpdate(context.Context, inputdata.UpdatePerson) error
}