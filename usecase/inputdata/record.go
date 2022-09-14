package inputdata

type AddRecord struct {
	MusicID      int
	RecordEasy   int
	ScoreEasy    int
	RecordNormal int
	ScoreNormal  int
	RecordHard   int
	ScoreHard    int
	RecordExpert int
	ScoreExpert  int
	RecordMaster int
	ScoreMaster  int
}

// The musicID to be updated is included in request URI.
type UpdateRecord struct {
	RecordEasy   int
	ScoreEasy    int
	RecordNormal int
	ScoreNormal  int
	RecordHard   int
	ScoreHard    int
	RecordExpert int
	ScoreExpert  int
	RecordMaster int
	ScoreMaster  int
}
