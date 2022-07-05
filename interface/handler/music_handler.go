package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/domain/model"
)

type MusicHandler struct {
	musicLogic model.MusicLogic
}

func NewMusicHandler(m model.MusicLogic) *MusicHandler {
	return &MusicHandler{
		musicLogic: m,
	}
}

func (h *MusicHandler) Get(ctx HttpContext) {
	musics, err := h.musicLogic.Fetch()
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("楽曲情報一覧が取得できません。"))
		return
	}

	output, _ := json.Marshal(musics)
	ctx.Response(http.StatusOK, output)
}
