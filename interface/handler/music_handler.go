package handler

import (
	"encoding/json"
	"fmt"
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

func (handler *MusicHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	musicList, err := handler.logic.GetMusicList()
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Fprintf(w, "Error: %s", err)
		return
	}

	output, _ := json.Marshal(musicList)
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
