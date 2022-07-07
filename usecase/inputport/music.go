package inputport

import "sekareco_srv/domain/model"

type MusicInputport interface {
	Fetch() ([]model.Music, error)
}
