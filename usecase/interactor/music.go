package interactor

import (
	"context"
	_infra "sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"

	"github.com/pkg/errors"
)

type musicInteractor struct {
	music       database.MusicRepository
	transaction database.SqlTransaction
}

func NewMusicInteractor(m database.MusicRepository, tx database.SqlTransaction) inputport.MusicInputport {
	return &musicInteractor{
		music:       m,
		transaction: tx,
	}
}

func (l *musicInteractor) Fetch(ctx context.Context) (musics []outputdata.Music, err error) {
	if musics, err = l.music.Fetch(ctx); err != nil {
		_infra.Logger.Error(errors.Wrap(err, "failed to select music"))
	}
	return
}

// interface implemention checks
var _ inputport.MusicInputport = &musicInteractor{}
