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

func (h *musicHandler) Get(ctx context.Context, hc infra.HttpContext) {
	musics, err := h.music.Fetch(ctx)
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("楽曲情報一覧が取得できません。"))
		return
	}

	output, _ := json.Marshal(musics)
	hc.Response(http.StatusOK, output)
}
