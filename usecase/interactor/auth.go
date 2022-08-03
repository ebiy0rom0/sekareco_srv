package interactor

import (
	"context"
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
	return i.token.GenerateNewToken()
}

func (i *authInteractor) AddToken(id int, token string) {
	i.token.AddToken(id, token)
}

func (i *authInteractor) RevokeToken(id int) {
	i.token.RevokeToken(id)
}

// interface implemention checks
var _ inputport.AuthInputport = &authInteractor{}
