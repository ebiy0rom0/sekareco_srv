package database

import (
	"context"
	"sekareco_srv/usecase/outputdata"
)

type MusicRepository interface {
	Fetch(context.Context) ([]outputdata.Music, error)
}
