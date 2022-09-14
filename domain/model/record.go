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
	ScoreEasy    int `json:"score_easy"`
	RecordNormal int `json:"record_normal"`
	ScoreNormal  int `json:"score_normal"`
	RecordHard   int `json:"record_hard"`
	ScoreHard    int `json:"score_hard"`
	RecordExpert int `json:"record_expert"`
	ScoreExpert  int `json:"score_expert"`
	RecordMaster int `json:"record_master"`
	ScoreMaster  int `json:"score_master"`
}
