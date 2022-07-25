package handler

import (
	"context"
	"encoding/json"
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

// @Summary		wip
// @Description	get all music master records
// @Tags		musics
// @Accept		json
// @Produce		json
// @Success		200
// @Failure		503
// @Router		/musics	[get]
// @SecurityDefinitions.apikey	Authentication
// @in							header
func (h *musicHandler) Get(ctx context.Context, hc infra.HttpContext) {
	musics, err := h.music.Fetch(ctx)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError(err))
		return
	}

	output, _ := json.Marshal(musics)
	hc.Response(http.StatusOK, output)
}
