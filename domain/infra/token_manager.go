package infra

type TokenManager interface {
	AddToken(int, string)
	RevokeToken(int)
	GenerateNewToken() string
}
