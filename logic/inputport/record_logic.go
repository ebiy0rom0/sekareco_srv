package inputport

import (
	"context"
	"sekareco_srv/domain/model"
)

type RecordLogic interface {
	Store(context.Context, model.Record) (int, error)
	Update(int, int, model.Record) error
	GetByPersonID(int) ([]model.Record, error)
}
