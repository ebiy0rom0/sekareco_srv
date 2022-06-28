package auth

import "golang.org/x/crypto/bcrypt"

type AuthLogic struct {
	Repository AuthRepository
}

func (logic *AuthLogic) CheckAuth(loginID string, password string) (personID int, err error) {
	login, err := logic.Repository.GetLoginPerson(loginID)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(login.PasswordHash), []byte(password))
	if err != nil {
		return
	}

	personID = login.PersonID
	return
}
