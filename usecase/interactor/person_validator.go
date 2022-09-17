package interactor

import (
	"context"
	"database/sql"
	"errors"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
)

type personValidator struct {
	login  database.LoginRepository
}

func NewPersonValidator(l database.LoginRepository) *personValidator {
	return &personValidator{
		login: l,
	}
}

func (v *personValidator) ValidationAdd(ctx context.Context, p inputdata.AddPerson) error {
	if ng := v.tooShortLoginID(p.LoginID); ng {
		return errors.New("too short loginID, need length upper 4 words")
	}
	if ng, err := v.duplicateLoginID(ctx, p.LoginID); err != nil {
		return err
	} else if ng {
		return errors.New("duplicate loginID")
	}
	if ng := v.tooShortPassword(p.Password); ng {
		return errors.New("too short password. need length upper 8 words")
	}
	if ng := v.requirePersonName(&p.PersonName); ng {
		return errors.New("person name is require")
	}

	return nil
}

func (v *personValidator) ValidationUpdate(ctx context.Context, p inputdata.UpdatePerson) error {
	return nil
}

func (v *personValidator) tooShortLoginID(loginID string) bool {
	return len(loginID) <= 4
}

func (v *personValidator) duplicateLoginID(ctx context.Context, loginID string) (bool, error) {
	if _, err := v.login.GetByID(ctx, loginID); err == sql.ErrNoRows {
		return true, nil
	} else if err != nil {
		return false, err
	}
	return false, nil
}

func (v *personValidator) requirePersonName(name *string) bool {
	return name == nil
}

func (v *personValidator) tooShortPassword(password string) bool {
	return len(password) <= 8
}

var _ inputport.PersonValidator = (*personValidator)(nil)