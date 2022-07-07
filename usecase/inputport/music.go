package inputport

import (
	"context"
	"sekareco_srv/domain/model"
)

type MusicInputport interface {
	Fetch(context.Context) ([]model.Music, error)
}
