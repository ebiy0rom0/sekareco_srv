package inputport

type AuthInputport interface {
	CheckAuth(string, string) (int, error)
	GenerateNewToken() string
}
