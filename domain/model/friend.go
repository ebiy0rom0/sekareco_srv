package model

type Friend struct {
	FriendID   int    `db:"friend_id" json:"friend_id"`
	PersonID   int    `db:"person_id" json:"person_id"`
	PersonName string `db:"person_name" json:"person_name"`
}
