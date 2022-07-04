package model

const (
	RECORD_NO_PLAY = iota
	RECORD_CLEAR
	RECORD_FULL_COMBO
	RECORD_ALL_PERFECT
)

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

type RecordRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	Store(Record) (int, error)
	Update(int, int, Record) error
	GetByPersonID(int) ([]Record, error)
}

type RecordLogic interface {
	Store(Record) (int, error)
	Update(int, int, Record) error
	GetByPersonID(int) ([]Record, error)
}
