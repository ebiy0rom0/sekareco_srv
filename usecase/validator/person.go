package validator

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

// ValidationAdd is the validation at the time of register new person.
func (v *personValidator) ValidationAdd(ctx context.Context, p inputdata.AddPerson) error {
	if ng := v.tooShortLoginID(p.LoginID); ng {
		return errors.New("too short loginID, need length 4 characters or over")
	}
	if ng, err := v.duplicateLoginID(ctx, p.LoginID); err != nil {
		return err
	} else if ng {
		return errors.New("duplicate loginID")
	}
	if ng := v.tooShortPassword(p.Password); ng {
		return errors.New("too short password. need length 8 characters or over")
	}
	if ng := v.requirePersonName(p.PersonName); ng {
		return errors.New("person name is require")
	}

	return nil
}

// ValidationUpdate is the validation at the time of update registered person info.
func (v *personValidator) ValidationUpdate(ctx context.Context, p inputdata.UpdatePerson) error {
	return nil
}

// tooShortLoginID checks to length of loginID.
// Returns true if less than 4 characters.
func (v *personValidator) tooShortLoginID(loginID string) bool {
	return len(loginID) < 4
}

// tooShortLoginID checks to duplicate loginID.
// Returns true if a duplicate loginID exists.
func (v *personValidator) duplicateLoginID(ctx context.Context, loginID string) (bool, error) {
	if _, err := v.login.GetByID(ctx, loginID); err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return true, err
	}
	return true, nil
}

// requirePersonName checks person name has been entered.
// Returns true if not entered person name.
func (v *personValidator) requirePersonName(name string) bool {
	return len(name) == 0
}

// tooShortPassword checks to length of password.
// Returns true if less than 8 characters.
func (v *personValidator) tooShortPassword(password string) bool {
	return len(password) < 8
}

var _ inputport.PersonValidator = (*personValidator)(nil)