package music

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
)

//	@Summary		get list | get all music master records
//	@Description	get all music master records
//	@Tags			music
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.Music
//	@Failure		503	{object}	infra.HttpError
//	@Security		Bearer Authentication
//	@Router			/musics	[get]
func (h *handler) Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError {
	musics, err := h.musicInputport.Fetch(ctx)
	if err != nil {
		return &infra.HttpError{Msg: err.Error(), Code: http.StatusServiceUnavailable}
	}

	hc.Response(http.StatusOK, musics)
	return nil
}
