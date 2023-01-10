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

func NewLoginRepository(h infra.SqlHandler) *loginRepository {
	return &loginRepository{h}
}

func (r *loginRepository) Store(ctx context.Context, l model.Login) error {
	query := "INSERT INTO person_login (login_id, person_id, password_hash)"
	query += " VALUES (?, ?, ?);"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	if _, err := dao.Execute(ctx, query, l.LoginID, l.PersonID, l.PasswordHash); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *loginRepository) GetByID(ctx context.Context, loginID string) (model.Login, error) {
	query := "SELECT password_hash, person_id FROM person_login WHERE login_id = ?;"
	row := r.QueryRow(ctx, query, loginID)

	var (
		personID     int
		passwordHash string
	)
	if err := row.Scan(&passwordHash, &personID); err != nil {
		return model.Login{}, errors.WithStack(err)
	}

	login := model.Login{
		PasswordHash: passwordHash,
		PersonID:     personID,
	}
	return login, nil
}

var _ database.LoginRepository = (*loginRepository)(nil)
