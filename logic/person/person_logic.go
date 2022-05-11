package person

import (
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"

	"golang.org/x/crypto/bcrypt"
)

type PersonLogic struct {
	Repository PersonRepository
}

func (logic *PersonLogic) RegistPerson(p model.Person) (personId int, err error) {
	personId, err = logic.Repository.RegistPerson(p)
	return
}

func (logic *PersonLogic) RegistLogin(l model.Login) (err error) {
	err = logic.Repository.RegistLogin(l)
	return
}

func (logic *PersonLogic) GetPersonById(personId int) (person model.Person, err error) {
	person, err = logic.Repository.GetPersonById(personId)
	return
}

func (logic *PersonLogic) CheckDuplicateLoginId(loginId string) (bool, error) {
	var ok bool
	person, err := logic.Repository.GetLoginPerson(loginId)
	if err != nil && err != sql.ErrNoRows {
		return ok, err
	}

	ok = person.PersonId == 0
	return ok, nil
}

func (logic *PersonLogic) GenerateFriendCode(loginId string) (code int) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	code, _ = fnv.New32().Write([]byte(loginId))
	return
}

func (logic *PersonLogic) GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
