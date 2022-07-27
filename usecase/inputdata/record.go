package inputdata

type AddRecord struct {
	MusicID      int
	RecordEasy   int
	RecordNormal int
	RecordHard   int
	RecordExpert int
	RecordMaster int
}

type UpdateRecord struct {
	RecordEasy   int
	RecordNormal int
	RecordHard   int
	RecordExpert int
	RecordMaster int
}
