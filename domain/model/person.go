package model

type Person struct {
	PersonID   int    `db:"person_id" json:"person_id"`
	PersonName string `db:"person_name" json:"person_name"`
	FriendCode int    `db:"friend_code" json:"friend_code"`
	IsCompare  bool   `db:"is_compare" json:"is_compare"`
}
