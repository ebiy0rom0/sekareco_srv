package logic

import (
	"sekareco_srv/domain/model"
	_infra "sekareco_srv/infra"
	"sekareco_srv/logic/database"
	"sekareco_srv/logic/inputport"

	"github.com/pkg/errors"
)

type MusicLogic struct {
	musicRepo   database.MusicRepository
	transaction database.SqlTransaction
}

func NewMusicLogic(m database.MusicRepository, tx database.SqlTransaction) inputport.MusicLogic {
	return &MusicLogic{
		musicRepo:   m,
		transaction: tx,
	}
}

func (l *MusicLogic) Fetch() (musics []model.Music, err error) {
	if musics, err = l.musicRepo.Fetch(); err != nil {
		_infra.Logger.Error(errors.Wrap(err, "failed to select music"))
	}
	return
}
