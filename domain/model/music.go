package model

type Music struct {
	MusicId     int    `json:"music_id"`
	ArtistId    int    `json:"artist_id"`
	MusicName   string `json:"music_name"`
	JacketUrl   string `json:"jacket_url"`
	LevelEasy   int    `json:"level_easy"`
	LevelNormal int    `json:"level_normal"`
	LevelHard   int    `json:"level_hard"`
	LevelExpert int    `json:"level_expert"`
	LevelMaster int    `json:"level_master"`
}

type MusicList []Music

type Artist struct {
	ArtistId   int    `json:"artist_id"`
	ArtistName string `json:"artist_name"`
	LogoUrl    string `json:"logo_url"`
}

type ArtistList []Artist
