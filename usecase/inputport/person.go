package inputport

import (
	"context"
	"sekareco_srv/domain/model"
)

type PersonInputport interface {
	Store(context.Context, model.PostPerson) (model.Person, error)
	GetByID(context.Context, int) (model.Person, error)
	IsDuplicateLoginID(context.Context, string) (bool, error)
}
