package outputdata

type Music struct {
	MusicID   int    `json:"musicID"`
	ArtistID  int    `json:"artistID"`
	MusicName string `json:"musicName"`
	JacketURL string `json:"jacketUrl"`
	Level     []int  `json:"level"`
}
