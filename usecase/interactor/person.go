package interactor

import (
	"context"
	"hash/fnv"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"

	"github.com/ebiy0rom0/errors"
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
			return nil, errors.New(err.Error())
		}
		person.PersonID = personID

		hash, err := i.toHashPassword(p.Password)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		login := model.Login{
			LoginID:      p.LoginID,
			PersonID:     personID,
			PasswordHash: hash,
		}

		if err = i.login.Store(ctx, login); err != nil {
			return nil, errors.New(err.Error())
		}
		return nil, nil

	}); err != nil {
		return model.Person{}, errors.New(err.Error())
	}
	return person, nil
}

func (i *personInteractor) Update(ctx context.Context, pid int, p inputdata.UpdatePerson) error {
	// TODO: create update in repository
	return nil
}

func (i *personInteractor) GetByID(ctx context.Context, personID int) (model.Person, error) {
	person, err := i.person.GetByID(ctx, personID)
	if err != nil {
		return model.Person{}, errors.New(err.Error())
	}
	return person, nil
}

func (i *personInteractor) generateFriendCode(loginID string) (int, error) {
	// Failed generate is not problem now.
	// This parameter usage in future content.
	h := fnv.New32()
	if _, err := h.Write([]byte(loginID)); err != nil {
		return 0, errors.New(err.Error())
	}

	return int(h.Sum32()), nil
}

func (i *personInteractor) toHashPassword(password string) (string, error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string(bhash), nil
}

var _ inputport.PersonInputport = (*personInteractor)(nil)
