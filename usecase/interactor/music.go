package interactor

import (
	"sekareco_srv/domain/model"
	_infra "sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"github.com/pkg/errors"
)

type MusicInteractor struct {
	musicRepo   database.MusicRepository
	transaction database.SqlTransaction
}

func NewMusicInteractor(m database.MusicRepository, tx database.SqlTransaction) inputport.MusicInputport {
	return &MusicInteractor{
		musicRepo:   m,
		transaction: tx,
	}
}

func (l *MusicInteractor) Fetch() (musics []model.Music, err error) {
	if musics, err = l.musicRepo.Fetch(); err != nil {
		_infra.Logger.Error(errors.Wrap(err, "failed to select music"))
	}
	return
}
