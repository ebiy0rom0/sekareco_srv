package model

type Login struct {
	LoginID      string `db:"login_id" json:"login_id"`
	PasswordHash string `db:"password_hash" json:"password_hash"`
	PersonID     int    `db:"person_id" json:"person_id"`
}
