package outputdata

type Record struct {
	MusicID int   `json:"music_id"`
	Records []int `json:"records"`
	Scores  []int `json:"scores"`
}