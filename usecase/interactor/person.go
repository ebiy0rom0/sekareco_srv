package interactor

import (
	"context"
	"fmt"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"strconv"

	"github.com/ebiy0rom0/errors"
	"golang.org/x/crypto/bcrypt"
)

type personInteractor struct {
	person      database.PersonRepository
	login       database.LoginRepository
	transaction database.SqlTransaction
}

func NewPersonInteractor(p database.PersonRepository, l database.LoginRepository, tx database.SqlTransaction) *personInteractor {
	return &personInteractor{
		person:      p,
		login:       l,
		transaction: tx,
	}
}

func (i *personInteractor) Store(ctx context.Context, p inputdata.AddPerson) (model.Person, error) {
	person := model.Person{
		PersonName: p.PersonName,
		IsCompare:  false,
	}

	if _, err := i.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		personID, err := i.person.Store(ctx, person)
		if err != nil {
			return nil, errors.New(err.Error())
		}
		person.PersonID = personID
		person.FriendCode = i.generateFriendCode(personID)

		if err := i.person.AddFriendCode(ctx, person); err != nil {
			return nil, errors.New(err.Error())
		}

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

func (i *personInteractor) generateFriendCode(personID int) int {
	var codeConvertMap = [][]uint8{
		{5, 1, 3, 6, 8, 9, 2, 4, 7, 0},
		{0, 3, 6, 9, 1, 7, 4, 5, 8, 2},
		{2, 7, 8, 9, 0, 4, 1, 6, 5, 3},
		{8, 6, 9, 0, 1, 5, 7, 2, 3, 4},
		{6, 4, 3, 9, 7, 8, 5, 1, 2, 0},
		{4, 0, 7, 1, 3, 8, 6, 9, 5, 2},
		{0, 1, 2, 3, 8, 6, 4, 5, 9, 7},
		{0, 4, 1, 3, 8, 6, 9, 7, 5, 2},
	}

	sID := fmt.Sprintf("%8d", personID)
	var friendCode int
	for n, r := range sID {
		digit, _ := strconv.Atoi(string(r))
		friendCode += friendCode*10 + int(codeConvertMap[n][digit])
	}
	return friendCode
}

func (i *personInteractor) toHashPassword(password string) (string, error) {
	bhash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return string(bhash), nil
}

var _ inputport.PersonInputport = (*personInteractor)(nil)
