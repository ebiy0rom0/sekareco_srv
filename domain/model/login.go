package model

type Login struct {
	LoginID      string `json:"login_id"`
	PasswordHash string `json:"password_hash"`
	PersonID     int    `json:"person_id"`
}

type LoginRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	Store(Login) error
	GetByID(string) (Login, error)
}
