package inputport

import "context"

type AuthInputport interface {
	CheckAuth(context.Context, string, string) (int, error)
	GenerateNewToken() string
}
