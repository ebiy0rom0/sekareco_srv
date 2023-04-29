package health

import (
	"context"
	"sekareco_srv/interface/infra"
)

type handler struct {
}

type Handler interface {
	Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError
}

func NewHealthHandler() *handler {
	return &handler{}
}
