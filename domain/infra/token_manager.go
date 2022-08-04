package infra

import "context"

var tokenKey = struct{}{}

type TokenManager interface {
	AddToken(int, string)
	RevokeToken(string)
	GenerateNewToken() string
}

func SetToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, &tokenKey, token)
}

func GetToken(ctx context.Context) string {
	return ctx.Value(&tokenKey).(string)
}
