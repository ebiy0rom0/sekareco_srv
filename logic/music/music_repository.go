package music

import "sekareco_srv/domain/model"

type MusicRepository interface {
	SelectAll() (model.MusicList, error)
}
