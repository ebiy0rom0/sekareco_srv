package model

type Person struct {
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	FriendCode int    `json:"friend_code"`
}

type Login struct {
	LoginID      string `json:"login_id"`
	PasswordHash string `json:"password_hash"`
	PersonID     int    `json:"person_id"`
}
