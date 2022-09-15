package interactor

import (
	"context"
	"database/sql"
	"hash/fnv"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type personInteractor struct {
	person      database.PersonRepository
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewPersonInteractor(p database.PersonRepository, l database.LoginRepository, tx database.SqlTransaction) inputport.PersonInputport {
	return &personInteractor{
		person:      p,
		login:       l,
		transaction: tx,
	}
}

func (i *personInteractor) Store(ctx context.Context, p inputdata.AddPerson) (model.Person, error) {
	code, _ := i.generateFriendCode(p.LoginID)
	person := model.Person{
		PersonName: p.PersonName,
		FriendCode: code,
	}

	if _, err := i.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		personID, err := i.person.Store(ctx, person)
		if err != nil {
			return nil, err
		}
		person.PersonID = personID

		hash, err := i.toHashPassword(p.Password)
		if err != nil {
			return nil, err
		}
		login := model.Login{
			LoginID:      p.LoginID,
			PersonID:     personID,
			PasswordHash: hash,
		}

		if err = i.login.Store(ctx, login); err != nil {
			return nil, err
		}
		return nil, nil

	}); err != nil {
		return model.Person{}, err
	}
	return person, nil
}

func (i *personInteractor) Update(ctx context.Context, pid int, p inputdata.UpdatePerson) error {
	// TODO: create update in repository
	return nil
}

func (i *personInteractor) GetByID(ctx context.Context, personID int) (person model.Person, err error) {
	if person, err = i.person.GetByID(ctx, personID); err != nil {
		return model.Person{}, errors.Wrapf(err, "failed to select person: person_id=%d", personID)
	}
	return
}

func (i *personInteractor) IsDuplicateLoginID(ctx context.Context, loginID string) (bool, error) {
	_, err := i.login.GetByID(ctx, loginID)
	if err == sql.ErrNoRows {
		return true, nil

	} else if err != nil {
		return false, errors.Wrapf(err, "failed to select login: login_id=%s", loginID)
	}

	return false, nil
}

func (i *personInteractor) generateFriendCode(loginID string) (code int, err error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	h := fnv.New32()
	if _, err := h.Write([]byte(loginID)); err != nil {
		infra.Logger.Warn(errors.Wrapf(err, "failed to generate friend code: login_id=%s", loginID))
	}

	code = int(h.Sum32())
	return
}

func (i *personInteractor) toHashPassword(password string) (hash string, err error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		infra.Logger.Error(errors.Wrap(err, "failed to generate password hash"))
		return
	}

	hash = string(bhash)
	return
}

var _ inputport.PersonInputport = (*personInteractor)(nil)
