package database

import "sekareco_srv/domain"

type PersonRepository struct {
	Handler SqlHandler
}

func (repository *PersonRepository) Regist(p domain.Person) (personId int, err error) {
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

func (repository *PersonRepository) SelectOne(personId int) (user domain.Person, err error) {
	rows := repository.Handler.QueryRow("SELECT person_id, person_name, firend_code FROM person WHERE person_id = ?", personId)

	var (
		personName string
		friendCode string
	)
	err = rows.Scan(&personId, &personName, &friendCode)
	if err != nil {
		return
	}

	user.PersonID = personId
	user.PersonName = personName
	user.FriendCode = friendCode
	return
}
