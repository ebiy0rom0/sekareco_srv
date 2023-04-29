package interactor

import (
	"context"
	"sekareco_srv/domain/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"github.com/ebiy0rom0/errors"

	"golang.org/x/crypto/bcrypt"
)

type authInteractor struct {
	token       infra.TokenManager
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewAuthInteractor(t infra.TokenManager, l database.LoginRepository, tx database.SqlTransaction) *authInteractor {
	return &authInteractor{
		token:       t,
		login:       l,
		transaction: tx,
	}
}

func (i *authInteractor) CheckAuth(ctx context.Context, loginID string, password string) (int, error) {
	login, err := i.login.GetByID(ctx, loginID)
	if err != nil {
		return 0, errors.Wrapf(err, "unregistered loginID: loginID=%s", loginID)
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(password))
	if err != nil {
		return 0, errors.Wrap(err, "password unmatch")
	}

	return login.PersonID, nil
}

func (i *authInteractor) AddToken(id int) infra.Token {
	token := i.token.GenerateNewToken()
	i.token.AddToken(id, token)
	return token
}

func (i *authInteractor) RevokeToken(token infra.Token) {
	i.token.RevokeToken(token)
}

var _ inputport.AuthInputport = (*authInteractor)(nil)
