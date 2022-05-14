package music

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"
)

type MusicLogic struct {
	Repository MusicRepository
}

func (logic *MusicLogic) GetMusicList() (musicList model.MusicList, err error) {
	if musicList, err = logic.Repository.SelectAll(); err != nil {
		logger.Logger.Error(err)
	}
	return
}
