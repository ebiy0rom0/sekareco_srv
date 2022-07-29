package inputdata

import (
	"errors"
)

type AddPerson struct {
	LoginID    string
	PersonName string
	Password   string
}

func (p *AddPerson) Valiation() error {
	if len(p.LoginID) == 0 {
		return errors.New("loginID is required")
	}
	if len(p.Password) == 0 {
		return errors.New("password is required")
	}
	if len(p.Password) < 8 {
		return errors.New("insufficient password security policy")
	}
	if len(p.PersonName) == 0 {
		return errors.New("person name is required")
	}
	return nil
}

type UpdatePerson struct {
	LoginID    string
	PersonName string
	Password   string
}
