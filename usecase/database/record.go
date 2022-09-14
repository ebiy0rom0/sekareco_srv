package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/outputdata"
)

type RecordRepository interface {
	Store(context.Context, model.Record) (int, error)
	Update(context.Context, int, int, model.Record) error
	GetByPersonID(context.Context, int) ([]outputdata.Record, error)
}
