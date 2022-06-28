package database

import (
	"sekareco_srv/domain/model"
)

type AuthRepository struct {
	Handler SqlHandler
}

func (repository *AuthRepository) GetLoginPerson(loginID string) (login model.Login, err error) {
	query := "SELECT password_hash, person_id FROM person_login WHERE login_id = ?;"
	row := repository.Handler.QueryRow(query, loginID)

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
