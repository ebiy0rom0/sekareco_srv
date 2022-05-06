package domain

type Music struct {
	MusicId     int
	ArtistId    int
	MusicName   string
	JacketUrl   string
	LevelEasy   int
	LevelNormal int
	LevelHard   int
	LevelExpert int
	LevelMaster int
}

type MusicList []Music

type Artist struct {
	ArtistId   int
	ArtistName string
	LogoUrl    string
}

type ArtistList []Artist
