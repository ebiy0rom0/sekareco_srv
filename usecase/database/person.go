package database

import (
	"context"
	"sekareco_srv/domain/model"
)

type PersonRepository interface {
	Store(context.Context, model.Person) (int, error)
	GetByID(context.Context, int) (model.Person, error)
	AddFriendCode(context.Context, model.Person) error
	GetByFriendCode(context.Context, int) (model.Person, error)
}
