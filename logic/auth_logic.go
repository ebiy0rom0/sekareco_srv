package logic

import (
	"encoding/base64"
	"net/http"
	"sekareco_srv/domain/model"
	_infra "sekareco_srv/infra"
	"sekareco_srv/logic/database"
	"sekareco_srv/logic/inputport"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var EXPIRED_IN = 1 * time.Hour
var MAX_TOKENS = 100

var EXPIRED_TOKEN_DELETE_SPAN = 15 * time.Minute

type personToken struct {
	token     string
	expiredIn time.Time
}

type AuthLogic struct {
	loginRepo   database.LoginRepository
	transaction database.SqlTransaction
	// access token mapping
	// key: personID, value: token
	tokens map[int]*personToken
}

func NewAuthLogic(l database.LoginRepository, tx database.SqlTransaction) inputport.AuthLogic {
	return &AuthLogic{
		loginRepo:   l,
		transaction: tx,
		tokens:      make(map[int]*personToken, MAX_TOKENS),
	}
}

func (l *AuthLogic) CheckAuth(loginID string, password string) (personID int, err error) {
	login, err := l.loginRepo.GetByID(loginID)
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
	return base64.StdEncoding.EncodeToString([]byte(_infra.Timer.NowDatetime()))
}

func (l *AuthLogic) AddToken(pid int, token string) {
	l.tokens[pid] = &personToken{
		token:     token,
		expiredIn: _infra.Timer.Add(EXPIRED_IN),
	}
}

func (l *AuthLogic) RevokeToken(pid int) {
	delete(l.tokens, pid)
}

func (l *AuthLogic) GetHeaderToken(r *http.Request) string {
	token := r.Header.Get(model.REQUEST_HEADER)
	return strings.Trim(strings.Replace(token, "Bearer", "", -1), " ")
}

func (l *AuthLogic) IsEnabledToken(pid int, token string) bool {
	access, ok := l.tokens[pid]

	// not exist token or token unmatch or token expired
	return !ok || token != access.token || _infra.Timer.Before(access.expiredIn)
}

func (l *AuthLogic) DeleteExpiredToken() {
	for k, t := range l.tokens {
		if _infra.Timer.Before(t.expiredIn) {
			l.RevokeToken(k)
		}
	}
}
