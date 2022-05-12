package dto

// POST /person/ request
type RequestRegistPerson struct {
	loginId    string
	password   string
	personName string
}

func NewRequestRegistPerson(loginId string, password string, personName string) (*RequestRegistPerson, error) {
	// TODO: parameter validation

	return &RequestRegistPerson{
		loginId:    loginId,
		password:   password,
		personName: personName,
	}, nil
}

func (vo *RequestRegistPerson) GetLoginId() string {
	return vo.loginId
}

func (vo *RequestRegistPerson) GetPassword() string {
	return vo.password
}

func (vo *RequestRegistPerson) GetPersonName() string {
	return vo.personName
}
