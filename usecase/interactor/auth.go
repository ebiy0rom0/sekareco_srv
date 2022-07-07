package interactor

import (
	"context"
	"encoding/base64"
	"sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"golang.org/x/crypto/bcrypt"
)

type authInteractor struct {
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewAuthInteractor(l database.LoginRepository, tx database.SqlTransaction) inputport.AuthInputport {
	return &authInteractor{
		login:       l,
		transaction: tx,
	}
}

func (i *authInteractor) CheckAuth(ctx context.Context, loginID string, password string) (personID int, err error) {
	login, err := i.login.GetByID(ctx, loginID)
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

func (i *authInteractor) GenerateNewToken() string {
	return base64.StdEncoding.EncodeToString([]byte(infra.Timer.NowDatetime()))
}
