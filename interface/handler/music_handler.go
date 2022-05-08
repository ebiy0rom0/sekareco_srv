package handler

import (
	"net/http"
	"sekareco_srv/interface/database"
	logic "sekareco_srv/logic/music"
)

type MusicHandler struct {
	logic logic.MusicLogic
}

func NewMusicHandler(sqlHandler database.SqlHandler) *MusicHandler {
	return &MusicHandler{
		logic: logic.MusicLogic{
			Repository: &database.MusicRepository{
				Handler: sqlHandler,
			},
		},
	}
}

func (handler *MusicHandler) Get(w http.ResponseWriter, r *http.Request) {

}
