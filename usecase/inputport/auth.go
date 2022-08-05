package inputport

import "context"

type AuthInputport interface {
	CheckAuth(context.Context, string, string) (int, error)
	AddToken(int) string
	RevokeToken(string)
}
