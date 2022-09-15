package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
)

type loginRepository struct {
	infra.SqlHandler
}

func NewLoginRepository(h infra.SqlHandler) *loginRepository {
	return &loginRepository{h}
}

func (r *loginRepository) Store(ctx context.Context, l model.Login) (err error) {
	query := "INSERT INTO person_login (login_id, person_id, password_hash)"
	query += " VALUES (?, ?, ?);"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	_, err = dao.Execute(ctx, query, l.LoginID, l.PersonID, l.PasswordHash)
	return
}

func (r *loginRepository) GetByID(ctx context.Context, loginID string) (login model.Login, err error) {
	query := "SELECT password_hash, person_id FROM person_login WHERE login_id = ?;"
	row := r.QueryRow(ctx, query, loginID)

	var (
		personID     int
		passwordHash string
	)
	err = row.Scan(&passwordHash, &personID)
	if err != nil {
		return
	}

	login.PasswordHash = passwordHash
	login.PersonID = personID
	return
}

var _ database.LoginRepository = (*loginRepository)(nil)
