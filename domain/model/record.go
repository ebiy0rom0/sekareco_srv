package model

const (
	RECORD_NO_PLAY = iota
	RECORD_CLEAR
	RECORD_FULL_COMBO
	RECORD_ALL_PERFECT
)

type Record struct {
	RecordID     int `db:"record_id" json:"record_id"`
	PersonID     int `db:"person_id" json:"person_id"`
	MusicID      int `db:"music_id" json:"music_id"`
	RecordEasy   int `db:"record_easy" json:"record_easy"`
	ScoreEasy    int `db:"score_easy" json:"score_easy"`
	RecordNormal int `db:"record_normal" json:"record_normal"`
	ScoreNormal  int `db:"score_normal" json:"score_normal"`
	RecordHard   int `db:"record_hard" json:"record_hard"`
	ScoreHard    int `db:"score_hard" json:"score_hard"`
	RecordExpert int `db:"record_expert" json:"record_expert"`
	ScoreExpert  int `db:"score_expert" json:"score_expert"`
	RecordMaster int `db:"record_master" json:"record_master"`
	ScoreMaster  int `db:"score_master" json:"score_master"`
}
