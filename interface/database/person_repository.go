package database

import "sekareco_srv/domain"

type PersonRepository struct {
	Handler SqlHandler
}

func (repository *PersonRepository) RegistPerson(p domain.Person) (personId int, err error) {
	// TODO: wip
	result, err := repository.Handler.Execute("INSERT INTO person VALUES ", p)
	if err != nil {
		return
	}

	newId64, err := result.LastInsertId()
	if err != nil {
		return
	}

	personId = int(newId64)
	return
}

func (repository *PersonRepository) RegistLogin(l domain.Login) (err error) {
	_, err = repository.Handler.Execute("INSERT INTO person_login VALUES ", l)
	return
}

func (repository *PersonRepository) GetPersonById(personId int) (user domain.Person, err error) {
	row := repository.Handler.QueryRow("SELECT person_id, person_name, friend_code FROM person WHERE person_id = ?", personId)

	var (
		personName string
		friendCode int
	)
	err = row.Scan(&personId, &personName, &friendCode)
	if err != nil {
		return
	}

	user.PersonId = personId
	user.PersonName = personName
	user.FriendCode = friendCode
	return
}

func (repository *PersonRepository) GetLoginPerson(loginId string) (login domain.Login, err error) {
	row := repository.Handler.QueryRow("SELECT password_hash, person_id FROM person_login WHERE login_id = ?", loginId)

	var (
		personId     int
		passwordHash string
	)
	err = row.Scan(&passwordHash, &personId)
	if err != nil {
		return
	}

	login.PasswordHash = passwordHash
	login.PersonId = personId
	return
}
