package handler

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

type healthHandler struct {
}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) Get(ctx context.Context, hc infra.HttpContext) {
	hc.Response(http.StatusOK, nil)
}
