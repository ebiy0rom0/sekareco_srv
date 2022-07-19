package inputdata

import (
	"errors"
)

type PostPerson struct {
	LoginID    string
	PersonName string
	Password   string
}

func (p *PostPerson) Valiation() error {
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

type PutPerson struct {
	LoginID    string
	PersonName string
	Password   string
}
