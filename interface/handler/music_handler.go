package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/interface/infra"
	"sekareco_srv/logic/inputport"
)

type MusicHandler struct {
	musicLogic inputport.MusicLogic
}

func NewMusicHandler(m inputport.MusicLogic) *MusicHandler {
	return &MusicHandler{
		musicLogic: m,
	}
}

func (h *MusicHandler) Get(ctx infra.HttpContext) {
	musics, err := h.musicLogic.Fetch()
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("楽曲情報一覧が取得できません。"))
		return
	}

	output, _ := json.Marshal(musics)
	ctx.Response(http.StatusOK, output)
}
