package inputport

import (
	"context"
	"sekareco_srv/domain/infra"
	"sekareco_srv/usecase/inputdata"
)

type AuthInputport interface {
	CheckAuth(context.Context, string, string) (int, error)
	AddToken(int) infra.Token
	RevokeToken(infra.Token)
}

type AuthValidator interface {
	ValidationPost(inputdata.PostAuth) error
}