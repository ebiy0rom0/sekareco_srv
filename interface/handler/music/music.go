package music

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
)

type handler struct {
	musicInputport inputport.MusicInputport
}

type Handler interface {
	Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError
}

func NewMusicHandler(musicInputport inputport.MusicInputport) *handler {
	return &handler{
		musicInputport: musicInputport,
	}
}
