package logic

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"github.com/pkg/errors"
)

type MusicLogic struct {
	musicRepo model.MusicRepository
}

func NewMusicLogic(m model.MusicRepository) model.MusicLogic {
	return &MusicLogic{
		musicRepo: m,
	}
}

func (l *MusicLogic) Fetch() (musics []model.Music, err error) {
	if musics, err = l.musicRepo.Fetch(); err != nil {
		logger.Logger.Error(errors.Wrap(err, "failed to select music"))
	}
	return
}
