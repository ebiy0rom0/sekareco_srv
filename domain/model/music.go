package model

type Music struct {
	MusicID     int    `db:"music_id" json:"music_id"`
	GroupID     int    `db:"group_id" json:"group_id"`
	MusicName   string `db:"music_name" json:"music_name"`
	JacketURL   string `db:"jacket_url" json:"jacket_url"`
	LevelEasy   int    `db:"level_easy" json:"level_easy"`
	NotesEasy   int    `db:"notes_easy" json:"notes_easy"`
	LevelNormal int    `db:"level_normal" json:"level_normal"`
	NotesNormal int    `db:"notes_normal" json:"notes_normal"`
	LevelHard   int    `db:"level_hard" json:"level_hard"`
	NotesHard   int    `db:"notes_hard" json:"notes_hard"`
	LevelExpert int    `db:"level_expert" json:"level_expert"`
	NotesExpert int    `db:"notes_expert" json:"notes_expert"`
	LevelMaster int    `db:"level_master" json:"level_master"`
	NotesMaster int    `db:"notes_master" json:"notes_master"`
}

type Group struct {
	GroupID   int    `db:"group_id" json:"group_id"`
	GroupName string `db:"group_name" json:"group_name"`
	LogoURL   string `db:"logo_url" json:"logo_url"`
}
