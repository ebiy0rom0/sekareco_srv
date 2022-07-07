package database

import "sekareco_srv/domain/model"

type MusicRepository interface {
	Fetch() ([]model.Music, error)
}
