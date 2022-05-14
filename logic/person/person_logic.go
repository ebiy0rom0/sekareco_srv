package person

import (
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"golang.org/x/crypto/bcrypt"
)

type PersonLogic struct {
	Repository PersonRepository
}

func (logic *PersonLogic) RegistPerson(p model.Person) (personId int, err error) {
	if personId, err = logic.Repository.RegistPerson(p); err != nil {
		logger.Logger.Error(err)
	}
	return
}

func (logic *PersonLogic) RegistLogin(l model.Login) (err error) {
	if err = logic.Repository.RegistLogin(l); err != nil {
		logger.Logger.Error(err)
	}
	return
}

func (logic *PersonLogic) GetPersonById(personId int) (person model.Person, err error) {
	if person, err = logic.Repository.GetPersonById(personId); err != nil {
		logger.Logger.Error(err)
	}
	return
}

func (logic *PersonLogic) CheckDuplicateLoginId(loginId string) (ok bool, err error) {
	person, err := logic.Repository.GetLoginPerson(loginId)
	if err != sql.ErrNoRows {
		logger.Logger.Warn(err)
		return
	} else if err != nil {
		logger.Logger.Error(err)
		return
	}

	ok = person.PersonId == 0
	return
}

func (logic *PersonLogic) GenerateFriendCode(loginId string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	if code, err = fnv.New32().Write([]byte(loginId)); err != nil {
		logger.Logger.Warn(err)
	}
	return
}

func (logic *PersonLogic) GeneratePasswordHash(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		logger.Logger.Error(err)
		return
	}

	hash = string(bhash)
	return
}
