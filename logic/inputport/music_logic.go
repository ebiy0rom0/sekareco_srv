package inputport

import "sekareco_srv/domain/model"

type MusicLogic interface {
	Fetch() ([]model.Music, error)
}
