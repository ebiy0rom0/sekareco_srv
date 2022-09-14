package model

type Music struct {
	MusicID     int    `json:"music_id"`
	ArtistID    int    `json:"artist_id"`
	MusicName   string `json:"music_name"`
	JacketURL   string `json:"jacket_url"`
	LevelEasy   int    `json:"level_easy"`
	NotesEasy   int    `json:"notes_easy"`
	LevelNormal int    `json:"level_normal"`
	NotesNormal int    `json:"notes_normal"`
	LevelHard   int    `json:"level_hard"`
	NotesHard   int    `json:"notes_hard"`
	LevelExpert int    `json:"level_expert"`
	NotesExpert int    `json:"notes_expert"`
	LevelMaster int    `json:"level_master"`
	NotesMaster int    `json:"notes_master"`
}

type Artist struct {
	ArtistID   int    `json:"artist_id"`
	ArtistName string `json:"artist_name"`
	LogoURL    string `json:"logo_url"`
}
