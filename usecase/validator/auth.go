package validator

import (
	"errors"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
)

type authValidator struct{}

func NewAuthValidator() *authValidator {
	return &authValidator{}
}

func (v *authValidator) ValidationPost(a inputdata.PostAuth) error {
	if ng := v.requireLoginID(a.LoginID); ng {
		return errors.New("loginID is require")
	}
	if ng := v.requirePassword(a.Password); ng {
		return errors.New("password is require")
	}

	return nil
}

// requirePersonName checks loginID has been entered.
// Returns true if not entered loginID.
func (v *authValidator) requireLoginID(loginID string) bool {
	return len(loginID) == 0
}

// requirePassword checks password has been entered.
// Returns true if not entered password.
func (v *authValidator) requirePassword(password string) bool {
	return len(password) == 0
}

var _ inputport.AuthValidator = (*authValidator)(nil)