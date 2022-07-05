package model

type Person struct {
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	FriendCode int    `json:"friend_code"`
}

type PostPerson struct {
	PersonName string `json:"person_name"`
	LoginID    string `json:"login_id"`
	Password   string `json:"password"`
}
