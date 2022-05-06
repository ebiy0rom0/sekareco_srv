package domain

type Record struct {
	RecordId     int
	PersonId     int
	MusicId      int
	RecordEasy   int
	RecordNormal int
	RecordHard   int
	RecordExpert int
	RecordMaster int
}

type RecordList []Record
