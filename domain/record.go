package domain

type Record struct {
	RecordId     int `json:"record_id"`
	PersonId     int `json:"person_id"`
	MusicId      int `json:"music_id"`
	RecordEasy   int `json:"record_easy"`
	RecordNormal int `json:"record_normal"`
	RecordHard   int `json:"record_hard"`
	RecordExpert int `json:"record_expert"`
	RecordMaster int `json:"record_master"`
}

type RecordList []Record
