package interactor

import (
	"context"
	"errors"
	infra_ "sekareco_srv/domain/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"golang.org/x/crypto/bcrypt"
)

type authInteractor struct {
	token       infra_.TokenManager
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewAuthInteractor(t infra_.TokenManager, l database.LoginRepository, tx database.SqlTransaction) inputport.AuthInputport {
	return &authInteractor{
		token:       t,
		login:       l,
		transaction: tx,
	}
}

func (i *authInteractor) CheckAuth(ctx context.Context, loginID string, password string) (int, error) {
	login, err := i.login.GetByID(ctx, loginID)
	if err != nil {
		return 0, errors.New("unregistered loginID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(password))
	if err != nil {
		return 0, errors.New("password unmatch")
	}

	return login.PersonID, nil
}

func (i *authInteractor) AddToken(id int) infra_.Token {
	token := i.token.GenerateNewToken()
	i.token.AddToken(id, token)
	return token
}

func (i *authInteractor) RevokeToken(token infra_.Token) {
	i.token.RevokeToken(token)
}

// interface implemention checks
var _ inputport.AuthInputport = &authInteractor{}
