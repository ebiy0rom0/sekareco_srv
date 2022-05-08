package logic

import "sekareco_srv/domain"

type MusicLogic struct {
	Repository MusicRepository
}

func (logic *MusicLogic) GetMusicList() (musicList domain.MusicList, err error) {
	musicList, err = logic.Repository.SelectAll()
	return
}
