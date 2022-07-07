package interactor

import (
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type PersonInteractor struct {
	person      database.PersonRepository
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewPersonInteractor(p database.PersonRepository, l database.LoginRepository, tx database.SqlTransaction) inputport.PersonInputport {
	return &PersonInteractor{
		person:      p,
		login:       l,
		transaction: tx,
	}
}

func (i *PersonInteractor) Store(p model.PostPerson) (model.Person, error) {
	// l.personRepo.StartTransaction()

	code, _ := i.generateFriendCode(p.LoginID)
	person := model.Person{
		PersonName: p.PersonName,
		FriendCode: code,
	}
	personID, err := i.person.Store(person)
	if err != nil {
		// l.loginRepo.Rollback()
		return model.Person{}, err
	}
	person.PersonID = personID

	hash, err := i.toHashPassword(p.Password)
	if err != nil {
		// l.loginRepo.Rollback()
		return model.Person{}, err
	}
	login := model.Login{
		LoginID:      p.LoginID,
		PersonID:     personID,
		PasswordHash: hash,
	}

	if err = i.login.Store(login); err != nil {
		// l.personRepo.Rollback()
		return model.Person{}, err
	}

	// l.personRepo.Commit()
	return person, nil
}

func (i *PersonInteractor) GetByID(personID int) (person model.Person, err error) {
	if person, err = i.person.GetByID(personID); err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to select person: person_id=%d", personID))
	}
	return
}

func (i *PersonInteractor) IsDuplicateLoginID(loginID string) (bool, error) {
	_, err := i.login.GetByID(loginID)
	if err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to select login: login_id=%s", loginID))
		return false, err
	}

	return false, nil
}

func (i *PersonInteractor) generateFriendCode(loginID string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	if code, err = fnv.New32().Write([]byte(loginID)); err != nil {
		infra.Logger.Warn(errors.Wrapf(err, "failed to generate friend code: login_id=%s", loginID))
	}
	return
}

func (i *PersonInteractor) toHashPassword(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		infra.Logger.Error(errors.Wrap(err, "failed to generate password hash"))
		return
	}

	hash = string(bhash)
	return
}
