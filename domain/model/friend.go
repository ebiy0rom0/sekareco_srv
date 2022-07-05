package model

type Friend struct {
	PersonID   int    `json:"person_id"`
	PersonName string `json:"person_name"`
	IsCompare  bool   `json:"is_compare"`
}
