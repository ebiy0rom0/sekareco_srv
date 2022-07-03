package auth

import (
	"encoding/base64"
	"net/http"
	"sekareco_srv/infra/timer"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var EXPIRED_IN = 1 * time.Hour
var MAX_TOKENS = 100

var EXPIRED_TOKEN_DELETE_SPAN = 15 * time.Minute

// access token mapping
// key: personID, value: token
var savedTokens = make(map[int]*personToken, MAX_TOKENS)

// RFC 6750 Bearer Token Conform
const (
	REQUEST_HEADER  = "Authorization"
	RESPONSE_HEADER = "WWW-Authenticate"
)

const (
	MESSAGE_OK            = "Bearer realm=\"\""
	MESSAGE_UNAUTHORIZED  = "Bearer realm=\"token_required\""
	MESSAGE_BAD_REQUEST   = "Bearer error=\"invalid_request\""
	MESSAGE_INVALID_TOKEN = "Bearer error=\"invalid_token\""
	MESSAGE_FORBIDDEN     = "Bearer error=\"insufficient_scope\""
)

type personToken struct {
	token     string
	expiredIn time.Time
}

type AuthLogic struct {
	Repository AuthRepository
}

func (l *AuthLogic) CheckAuth(loginID string, password string) (personID int, err error) {
	login, err := l.Repository.GetLoginPerson(loginID)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(password))
	if err != nil {
		return
	}

	personID = login.PersonID
	return
}

func (l *AuthLogic) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(timer.Timer.NowDatetime()))
}

func (l *AuthLogic) AddToken(pid int, token string) {
	savedTokens[pid] = &personToken{
		token:     token,
		expiredIn: timer.Timer.Add(EXPIRED_IN),
	}
}

func (l *AuthLogic) RevokeToken(pid int) {
	delete(savedTokens, pid)
}

func (l *AuthLogic) GetInHeaderToken(r *http.Request) string {
	token := r.Header.Get(REQUEST_HEADER)
	return strings.Trim(strings.Replace(token, "Bearer", "", -1), " ")
}

func (l *AuthLogic) IsEnabledToken(pid int, token string) bool {
	access, ok := savedTokens[pid]

	// not exist token or token unmatch or token expired
	return !ok || token != access.token || timer.Timer.Before(access.expiredIn)
}

func (l *AuthLogic) DeleteExpiredToken() {
	for k, t := range savedTokens {
		if timer.Timer.Before(t.expiredIn) {
			l.RevokeToken(k)

		}
	}
}
