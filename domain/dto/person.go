package dto

// POST /person/ request
type RequestRegistPerson struct {
	loginID    string
	password   string
	personName string
}

func NewRequestRegistPerson(loginID string, password string, personName string) (*RequestRegistPerson, error) {
	// TODO: parameter validation

	return &RequestRegistPerson{
		loginID:    loginID,
		password:   password,
		personName: personName,
	}, nil
}

func (vo *RequestRegistPerson) GetLoginID() string {
	return vo.loginID
}

func (vo *RequestRegistPerson) GetPassword() string {
	return vo.password
}

func (vo *RequestRegistPerson) GetPersonName() string {
	return vo.personName
}
