package database

import (
	"sekareco_srv/domain/model"

	"github.com/pkg/errors"
)

type PersonRepository struct {
	Handler SqlHandler
}

func (repository *PersonRepository) StartTransaction() error {
	return repository.Handler.StartTransaction()
}

func (repository *PersonRepository) Commit() error {
	return repository.Handler.Commit()
}

func (repository *PersonRepository) Rollback() error {
	return repository.Handler.Rollback()
}

func (repository *PersonRepository) RegistPerson(p model.Person) (personID int, err error) {
	query := "INSERT INTO person (person_name, friend_code)"
	query += " VALUES (?, ?);"

	result, err := repository.Handler.Execute(query, p.PersonName, p.FriendCode)
	if err != nil {
		err = errors.Wrap(err, "failed")
		return
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return
	}

	personID = int(newID64)
	return
}

func (repository *PersonRepository) RegistLogin(l model.Login) (err error) {
	query := "INSERT INTO person_login (login_id, person_id, password_hash)"
	query += " VALUES (?, ?, ?);"

	_, err = repository.Handler.Execute(query, l.LoginID, l.PersonID, l.PasswordHash)
	return
}

func (repository *PersonRepository) GetPersonByID(personID int) (user model.Person, err error) {
	query := "SELECT person_id, person_name, friend_code FROM person WHERE person_id = ?;"
	row := repository.Handler.QueryRow(query, personID)

	var (
		personName string
		friendCode int
	)
	err = row.Scan(&personID, &personName, &friendCode)
	if err != nil {
		return
	}

	user.PersonID = personID
	user.PersonName = personName
	user.FriendCode = friendCode
	return
}

func (repository *PersonRepository) GetLoginPerson(loginID string) (login model.Login, err error) {
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
