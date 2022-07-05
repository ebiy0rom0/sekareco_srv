package database

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/logic/database"
)

type LoginRepository struct {
	infra.SqlHandler
}

func NewLoginRepository(h infra.SqlHandler) database.LoginRepository {
	return &LoginRepository{h}
}

func (r *LoginRepository) Store(l model.Login) (err error) {
	query := "INSERT INTO person_login (login_id, person_id, password_hash)"
	query += " VALUES (?, ?, ?);"

	_, err = r.Execute(query, l.LoginID, l.PersonID, l.PasswordHash)
	return
}

func (r *LoginRepository) GetByID(loginID string) (login model.Login, err error) {
	query := "SELECT password_hash, person_id FROM person_login WHERE login_id = ?;"
	row := r.QueryRow(query, loginID)

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
