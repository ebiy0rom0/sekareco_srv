package outputdata

type Music struct {
	MusicID   int    `json:"music_id"`
	ArtistID  int    `json:"artist_id"`
	MusicName string `json:"music_name"`
	JacketURL string `json:"jacket_url"`
	Level     []int  `json:"level"`
}
