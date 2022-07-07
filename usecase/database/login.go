package database

import (
	"context"
	"sekareco_srv/domain/model"
)

type LoginRepository interface {
	Store(context.Context, model.Login) error
	GetByID(context.Context, string) (model.Login, error)
}
