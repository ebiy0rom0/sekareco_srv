package person

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
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

func NewPersonHandler(personInputport inputport.PersonInputport, personValidator inputport.PersonValidator) Handler {
	return &handler{
		personInputport: personInputport,
		personValidator: personValidator,
	}
}
