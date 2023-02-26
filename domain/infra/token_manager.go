package infra

import (
	"context"

	"github.com/ebiy0rom0/errors"
)

var tokenKey = struct{}{}

type Token string

type sessionInfo struct {
	personID int
	token    Token
}

type TokenManager interface {
	AddToken(int, Token)
	RevokeToken(Token)
	GenerateNewToken() Token
}

func NewSessionInfo(pid int, token Token) sessionInfo {
	return sessionInfo{
		personID: pid,
		token:    token,
	}
}

func SetToken(ctx context.Context, info sessionInfo) context.Context {
	return context.WithValue(ctx, &tokenKey, info)
}

func GetToken(ctx context.Context) (Token, error) {
	info, err := getSessionInfo(ctx)
	if err != nil {
		return "", err
	}
	return info.token, nil
}

func GetPersonID(ctx context.Context) (int, error) {
	info, err := getSessionInfo(ctx)
	if err != nil {
		return 0, err
	}
	return info.personID, nil
}

func getSessionInfo(ctx context.Context) (sessionInfo, error) {
	v := ctx.Value(&tokenKey)
	if v == nil {
		return sessionInfo{}, errors.New("session info not set in context")
	}

	info, ok := v.(sessionInfo)
	if !ok {
		return sessionInfo{}, errors.New("session info not set in context")
	}
	return info, nil
}
