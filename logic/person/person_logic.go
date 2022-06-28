package person

import (
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type PersonLogic struct {
	Repository PersonRepository
}

func (logic *PersonLogic) RegistPerson(p model.Person) (personID int, err error) {
	if personID, err = logic.Repository.RegistPerson(p); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist person: %#v", p))
	}
	return
}

func (logic *PersonLogic) RegistLogin(l model.Login) (err error) {
	if err = logic.Repository.RegistLogin(l); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist login: %#v", l))
	}
	return
}

func (logic *PersonLogic) GetPersonByID(personID int) (person model.Person, err error) {
	if person, err = logic.Repository.GetPersonByID(personID); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select person: person_id=%d", personID))
	}
	return
}

func (logic *PersonLogic) CheckDuplicateLoginID(loginID string) (bool, error) {
	_, err := logic.Repository.GetLoginPerson(loginID)
	if err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select login: login_id=%s", loginID))
		return false, err
	}

	return false, nil
}

func (logic *PersonLogic) GenerateFriendCode(loginID string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	if code, err = fnv.New32().Write([]byte(loginID)); err != nil {
		logger.Logger.Warn(errors.Wrapf(err, "failed to generate friend code: login_id=%s", loginID))
	}
	return
}

func (logic *PersonLogic) GeneratePasswordHash(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		logger.Logger.Error(errors.Wrap(err, "failed to generate password hash"))
		return
	}

	hash = string(bhash)
	return
}
