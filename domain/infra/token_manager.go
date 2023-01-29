package infra

import "context"

var tokenKey = struct{}{}

type Token string

type TokenManager interface {
	AddToken(int, Token)
	RevokeToken(Token)
	GenerateNewToken() Token
}

func SetToken(ctx context.Context, token Token) context.Context {
	return context.WithValue(ctx, &tokenKey, token)
}

func GetToken(ctx context.Context) Token {
	v := ctx.Value(&tokenKey)
	if v == nil {
		return ""
	}
	return v.(Token)
}
