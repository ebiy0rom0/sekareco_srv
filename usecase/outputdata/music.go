package outputdata

type Music struct {
	MusicID   int    `json:"musicID"`
	GroupID   int    `json:"groupID"`
	MusicName string `json:"musicName"`
	JacketURL string `json:"jacketUrl"`
	Level     []int  `json:"level"`
	Notes     []int  `json:"notes"`
}
