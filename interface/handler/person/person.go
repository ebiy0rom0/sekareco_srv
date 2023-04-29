package person

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"

	"github.com/google/wire"
)

type handler struct {
	personInputport inputport.PersonInputport
	personValidator inputport.PersonValidator
}

type Handler interface {
	Get(ctx context.Context, hc infra.HttpContext) *infra.HttpError
	Post(ctx context.Context, hc infra.HttpContext) *infra.HttpError
	Put(ctx context.Context, hc infra.HttpContext) *infra.HttpError
}

func NewPersonHandler(personInputport inputport.PersonInputport, personValidator inputport.PersonValidator) *handler {
	return &handler{
		personInputport: personInputport,
		personValidator: personValidator,
	}
}

var PersonHandlerProviderSet = wire.NewSet(
	NewPersonHandler,
	wire.Bind(new(Handler), new(*handler)),
)

var _ Handler = (*handler)(nil)
