package health

import (
	"context"
	"sekareco_srv/interface/infra"

	"github.com/google/wire"
)

type handler struct {
}

type Handler interface {
	Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError
}

func NewHealthHandler() *handler {
	return &handler{}
}

var HealthHandlerProviderSet = wire.NewSet(
	NewHealthHandler,
	wire.Bind(new(Handler), new(*handler)),
)

var _ Handler = (*handler)(nil)
