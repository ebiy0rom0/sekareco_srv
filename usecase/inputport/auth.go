package inputport

import (
	"context"
	"sekareco_srv/domain/infra"
)

type AuthInputport interface {
	CheckAuth(context.Context, string, string) (int, error)
	AddToken(int) infra.Token
	RevokeToken(infra.Token)
}
