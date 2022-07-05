package database

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/logic/database"

	"github.com/pkg/errors"
)

type PersonRepository struct {
	infra.SqlHandler
}

func NewPersonRepository(h infra.SqlHandler) database.PersonRepository {
	return &PersonRepository{h}
}

func (r *PersonRepository) Store(p model.Person) (personID int, err error) {
	query := "INSERT INTO person (person_name, friend_code)"
	query += " VALUES (?, ?);"

	result, err := r.Execute(query, p.PersonName, p.FriendCode)
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

func (r *PersonRepository) GetByID(personID int) (user model.Person, err error) {
	query := "SELECT person_id, person_name, friend_code FROM person WHERE person_id = ?;"
	row := r.QueryRow(query, personID)

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
