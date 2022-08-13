package inputport

import (
	"context"
	"sekareco_srv/usecase/outputdata"
)

type MusicInputport interface {
	Fetch(context.Context) ([]outputdata.Music, error)
}
