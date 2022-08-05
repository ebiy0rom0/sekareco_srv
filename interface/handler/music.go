package handler

import (
	"context"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
)

type musicHandler struct {
	music inputport.MusicInputport
}

func NewMusicHandler(m inputport.MusicInputport) *musicHandler {
	return &musicHandler{
		music: m,
	}
}

// @Summary		get list | get all music master records
// @Description	get all music master records
// @Tags		music
// @Accept		json
// @Produce		json
// @Success		200	{object}	[]model.Music
// @Failure		503	{object}	infra.HttpError
// @Security	Bearer Authentication
// @Router		/musics	[get]
func (h *musicHandler) Get(ctx context.Context, hc infra.HttpContext) {
	musics, err := h.music.Fetch(ctx)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	hc.Response(http.StatusOK, musics)
}
