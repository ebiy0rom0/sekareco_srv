package handler

import (
	"encoding/json"
	"net/http"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/music"
)

type MusicHandler struct {
	logic music.MusicLogic
}

func NewMusicHandler(sqlHandler database.SqlHandler) *MusicHandler {
	return &MusicHandler{
		logic: music.MusicLogic{
			Repository: &database.MusicRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *MusicHandler) Get(ctx HttpContext) {
	musicList, err := handler.logic.GetMusicList()
	if err != nil {
		ctx.Response(http.StatusServiceUnavailable, ctx.MakeError("楽曲情報一覧が取得できません。"))
		return
	}

	output, _ := json.Marshal(musicList)
	ctx.Response(http.StatusOK, output)
}
