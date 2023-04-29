package auth

import (
	"context"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/inputport"
)

type handler struct {
	authInputport inputport.AuthInputport
	authValidator inputport.AuthValidator
}

type Handler interface {
	Post(context.Context, infra.HttpContext) *infra.HttpError
	Delete(context.Context, infra.HttpContext) *infra.HttpError
}

func NewAuthHandler(authInputPort inputport.AuthInputport, authValidator inputport.AuthValidator) *handler {
	return &handler{
		authInputport: authInputPort,
		authValidator: authValidator,
	}
}
