package model

type Music struct {
	MusicID     int    `json:"music_id"`
	ArtistID    int    `json:"artist_id"`
	MusicName   string `json:"music_name"`
	JacketURL   string `json:"jacket_url"`
	LevelEasy   int    `json:"level_easy"`
	LevelNormal int    `json:"level_normal"`
	LevelHard   int    `json:"level_hard"`
	LevelExpert int    `json:"level_expert"`
	LevelMaster int    `json:"level_master"`
}

type Artist struct {
	ArtistID   int    `json:"artist_id"`
	ArtistName string `json:"artist_name"`
	LogoURL    string `json:"logo_url"`
}

type MusicRepository interface {
	Fetch() ([]Music, error)
}

type MusicLogic interface {
	Fetch() ([]Music, error)
}
