package music

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"github.com/pkg/errors"
)

type MusicLogic struct {
	Repository MusicRepository
}

func (l *MusicLogic) GetMusicList() (musicList model.MusicList, err error) {
	if musicList, err = l.Repository.SelectAll(); err != nil {
		logger.Logger.Error(errors.Wrap(err, "failed to select music"))
	}
	return
}
