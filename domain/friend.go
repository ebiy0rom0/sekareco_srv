package domain

type Friend struct {
	PersonId   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	IsCompare  bool   `json:"is_compare"`
}

type FriendList []Friend
