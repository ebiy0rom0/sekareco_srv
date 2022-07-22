package inputport

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
)

type RecordInputport interface {
	Store(context.Context, inputdata.PostRecord) (int, error)
	Update(context.Context, int, int, inputdata.PutRecord) error
	GetByPersonID(context.Context, int) ([]model.Record, error)
}
