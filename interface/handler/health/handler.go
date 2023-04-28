package health

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

func (h *handler) Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	hc.Response(http.StatusOK, nil)
	return nil
}

var _ Handler = (*handler)(nil)
