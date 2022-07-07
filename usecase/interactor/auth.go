package interactor

import (
	"encoding/base64"
	"sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewAuthInteractor(l database.LoginRepository, tx database.SqlTransaction) inputport.AuthInputport {
	return &AuthInteractor{
		login:       l,
		transaction: tx,
	}
}

func (i *AuthInteractor) CheckAuth(loginID string, password string) (personID int, err error) {
	// l.Do
	login, err := i.login.GetByID(loginID)
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

func (i *AuthInteractor) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(infra.Timer.NowDatetime()))
}
