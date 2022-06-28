package middleware

// access token mapping
// key: personID, value: token
var tokens map[int]string

func AddTokens(pid int, token string) {
	if tokens == nil {
		tokens = make(map[int]string)
	}
	tokens[pid] = token
}

func RemoveTokens(pid int) {
	delete(tokens, pid)
}

func IsEnableToken(pid int, token string) bool {
	registeredToken, ok := tokens[pid]

	// not exist token or token unmatch
	return !ok || token != registeredToken
}
