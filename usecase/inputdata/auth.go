package inputdata

import "errors"

type PostAuth struct {
	LoginID  string
	Password string
}

func (a *PostAuth) Validation() error {
	if len(a.LoginID) == 0 {
		return errors.New("loginID is required")
	}
	if len(a.Password) == 0 {
		return errors.New("password is required")
	}
	return nil
}

type DeleteAuth struct {
	PersonID string
}

func (a *DeleteAuth) Validation() error {
	if len(a.PersonID) == 0 {
		return errors.New("personID is required")
	}
	return nil
}
