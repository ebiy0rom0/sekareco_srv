package logic

import (
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"
	_infra "sekareco_srv/infra"
	"sekareco_srv/logic/database"
	"sekareco_srv/logic/inputport"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type PersonLogic struct {
	personRepo  database.PersonRepository
	loginRepo   database.LoginRepository
	transaction database.SqlTransaction
}

func NewPersonLogic(p database.PersonRepository, l database.LoginRepository, tx database.SqlTransaction) inputport.PersonLogic {
	return &PersonLogic{
		personRepo:  p,
		loginRepo:   l,
		transaction: tx,
	}
}

func (l *PersonLogic) Store(p model.PostPerson) (model.Person, error) {
	// l.personRepo.StartTransaction()

	code, _ := l.generateFriendCode(p.LoginID)
	person := model.Person{
		PersonName: p.PersonName,
		FriendCode: code,
	}
	personID, err := l.personRepo.Store(person)
	if err != nil {
		// l.loginRepo.Rollback()
		return model.Person{}, err
	}
	person.PersonID = personID

	hash, err := l.toHashPassword(p.Password)
	if err != nil {
		// l.loginRepo.Rollback()
		return model.Person{}, err
	}
	login := model.Login{
		LoginID:      p.LoginID,
		PersonID:     personID,
		PasswordHash: hash,
	}

	if err = l.loginRepo.Store(login); err != nil {
		// l.personRepo.Rollback()
		return model.Person{}, err
	}

	// l.personRepo.Commit()
	return person, nil
}

func (l *PersonLogic) GetByID(personID int) (person model.Person, err error) {
	if person, err = l.personRepo.GetByID(personID); err != nil {
		_infra.Logger.Error(errors.Wrapf(err, "failed to select person: person_id=%d", personID))
	}
	return
}

func (l *PersonLogic) IsDuplicateLoginID(loginID string) (bool, error) {
	_, err := l.loginRepo.GetByID(loginID)
	if err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		_infra.Logger.Error(errors.Wrapf(err, "failed to select login: login_id=%s", loginID))
		return false, err
	}

	return false, nil
}

func (l *PersonLogic) generateFriendCode(loginID string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	if code, err = fnv.New32().Write([]byte(loginID)); err != nil {
		_infra.Logger.Warn(errors.Wrapf(err, "failed to generate friend code: login_id=%s", loginID))
	}
	return
}

func (l *PersonLogic) toHashPassword(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		_infra.Logger.Error(errors.Wrap(err, "failed to generate password hash"))
		return
	}

	hash = string(bhash)
	return
}
