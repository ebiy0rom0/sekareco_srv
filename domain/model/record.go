package model

type Record struct {
	RecordID     int `json:"record_id"`
	PersonID     int `json:"person_id"`
	MusicID      int `json:"music_id"`
	RecordEasy   int `json:"record_easy"`
	RecordNormal int `json:"record_normal"`
	RecordHard   int `json:"record_hard"`
	RecordExpert int `json:"record_expert"`
	RecordMaster int `json:"record_master"`
}

type RecordList []Record
