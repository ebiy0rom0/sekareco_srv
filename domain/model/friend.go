package model

type Friend struct {
	FriendID   int    `json:"friend_id"`
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
}
