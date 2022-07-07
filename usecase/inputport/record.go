package inputport

import (
	"context"
	"sekareco_srv/domain/model"
)

type RecordInputport interface {
	Store(context.Context, model.Record) (int, error)
	Update(context.Context, int, int, model.Record) error
	GetByPersonID(context.Context, int) ([]model.Record, error)
}
