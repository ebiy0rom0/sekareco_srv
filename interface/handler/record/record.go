package record

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"

	"github.com/google/wire"
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

var RecordHandlerProviderSet = wire.NewSet(
	NewRecordHandler,
	wire.Bind(new(Handler), new(*handler)),
)

var _ Handler = (*handler)(nil)
