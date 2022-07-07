package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
)

type MusicHandler struct {
	music inputport.MusicInputport
}

func NewMusicHandler(m inputport.MusicInputport) *MusicHandler {
	return &MusicHandler{
		music: m,
	}
}

func (h *MusicHandler) Get(ctx context.Context, hc infra.HttpContext) {
	musics, err := h.music.Fetch()
	if err != nil {
		hc.Response(http.StatusServiceUnavailable, hc.MakeError("楽曲情報一覧が取得できません。"))
		return
	}

	output, _ := json.Marshal(musics)
	hc.Response(http.StatusOK, output)
}
