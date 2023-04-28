package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"

	"github.com/ebiy0rom0/errors"
)

type loginRepository struct {
	infra.SqlHandler
}

func NewLoginRepository(h infra.SqlHandler) database.LoginRepository {
	return &loginRepository{h}
}

func (r *loginRepository) Store(ctx context.Context, l model.Login) error {
	query := `
	INSERT INTO person_login (
		login_id, person_id, password_hash
	) VALUES (
		:login_id, :person_id, :password_hash
	);`

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	if _, err := dao.ExecNamedContext(ctx, query, l); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *loginRepository) GetByID(ctx context.Context, loginID string) (model.Login, error) {
	query := `SELECT * FROM person_login WHERE login_id = $1;`

	var login model.Login
	if err := r.GetContext(ctx, &login, query, loginID); err != nil {
		return model.Login{}, errors.WithStack(err)
	}
	return login, nil
}

var _ database.LoginRepository = (*loginRepository)(nil)
