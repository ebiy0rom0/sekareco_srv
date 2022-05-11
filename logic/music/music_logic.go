package music

import "sekareco_srv/domain/model"

type MusicLogic struct {
	Repository MusicRepository
}

func (logic *MusicLogic) GetMusicList() (musicList model.MusicList, err error) {
	musicList, err = logic.Repository.SelectAll()
	return
}
