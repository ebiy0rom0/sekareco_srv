package record

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
)

type handler struct {
	recordInputport inputport.RecordInputport
}

type Handler interface {
	Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError
	Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError
	Put(ctx context.Context, hc infra.HttpContext) *infra.HttpError
}

func NewRecordHandler(recordInputport inputport.RecordInputport) *handler {
	return &handler{
		recordInputport: recordInputport,
	}
}
