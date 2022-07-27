package inputport

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
)

type RecordInputport interface {
	Store(context.Context, int, inputdata.AddRecord) (model.Record, error)
	Update(context.Context, int, int, inputdata.UpdateRecord) error
	GetByPersonID(context.Context, int) ([]model.Record, error)
}
