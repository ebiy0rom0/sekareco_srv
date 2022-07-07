package database

import (
	"context"
	"sekareco_srv/domain/model"
)

type MusicRepository interface {
	Fetch(context.Context) ([]model.Music, error)
}
