package database

import (
	"sekareco_srv/domain/model"
)

type RecordRepository interface {
	Store(model.Record) (int, error)
	Update(int, int, model.Record) error
	GetByPersonID(int) ([]model.Record, error)
	// Do(context.Context, ExecFunc) (interface{}, error)
}
