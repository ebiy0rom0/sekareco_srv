package music

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"

	"github.com/google/wire"
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

var MusicHandlerProviderSet = wire.NewSet(
	NewMusicHandler,
	wire.Bind(new(Handler), new(*handler)),
)

var _ Handler = (*handler)(nil)
