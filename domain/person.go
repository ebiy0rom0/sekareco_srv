package domain

type Person struct {
	PersonId   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	FriendCode int    `json:"friend_code"`
}

type Login struct {
	LoginId      string `json:"login_id"`
	PasswordHash string `json:"password_hash"`
	PersonId     int    `json:"person_id"`
}
