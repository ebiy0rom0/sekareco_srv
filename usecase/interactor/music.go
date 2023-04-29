package interactor

import (
	"context"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"

	"github.com/ebiy0rom0/errors"
)

type musicInteractor struct {
	music       database.MusicRepository
	transaction database.SqlTransaction
}

func NewMusicInteractor(m database.MusicRepository, tx database.SqlTransaction) *musicInteractor {
	return &musicInteractor{
		music:       m,
		transaction: tx,
	}
}

func (l *musicInteractor) Fetch(ctx context.Context) ([]outputdata.Music, error) {
	musics, err := l.music.Fetch(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to select music")
	}
	return musics, nil
}

var _ inputport.MusicInputport = (*musicInteractor)(nil)
