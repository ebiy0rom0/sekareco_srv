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

func (l *PersonLogic) RegistPerson(p model.Person) (personID int, err error) {
	if personID, err = l.Repository.RegistPerson(p); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist person: %#v", p))
	}
	return
}

func (l *PersonLogic) RegistLogin(lo model.Login) (err error) {
	if err = l.Repository.RegistLogin(lo); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist login: %#v", l))
	}
	return
}

func (l *PersonLogic) GetPersonByID(personID int) (person model.Person, err error) {
	if person, err = l.Repository.GetPersonByID(personID); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select person: person_id=%d", personID))
	}
	return
}

func (l *PersonLogic) CheckDuplicateLoginID(loginID string) (bool, error) {
	_, err := l.Repository.GetLoginPerson(loginID)
	if err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select login: login_id=%s", loginID))
		return false, err
	}

	return false, nil
}

func (l *PersonLogic) GenerateFriendCode(loginID string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	if code, err = fnv.New32().Write([]byte(loginID)); err != nil {
		logger.Logger.Warn(errors.Wrapf(err, "failed to generate friend code: login_id=%s", loginID))
	}
	return
}

func (l *PersonLogic) GeneratePasswordHash(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		logger.Logger.Error(errors.Wrap(err, "failed to generate password hash"))
		return
	}

	hash = string(bhash)
	return
}
