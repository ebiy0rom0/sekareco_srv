package database_test

import (
	"context"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"
	"testing"

	"github.com/stretchr/testify/assert"
)

// initialize at TestMain() in login_test.go
var musicRepo database.MusicRepository

func TestMusicRepository_Fetch(t *testing.T) {
	ctx := context.Background()
	want := []outputdata.Music{
		{
			MusicID:   1,
			ArtistID:  1,
			MusicName: "test_music001",
			JacketURL: "jacket/m_001.png",
			Level:     []int{1, 2, 3, 4, 5},
			Notes:     []int{100, 200, 300, 400, 500},
		},
		{
			MusicID:   2,
			ArtistID:  2,
			MusicName: "test_music002",
			JacketURL: "jacket/m_002.png",
			Level:     []int{2, 3, 4, 5, 6},
			Notes:     []int{200, 300, 400, 500, 600},
		},
		{
			MusicID:   3,
			ArtistID:  1,
			MusicName: "test_music003",
			JacketURL: "jacket/m_003.png",
			Level:     []int{3, 4, 5, 6, 7},
			Notes:     []int{300, 400, 500, 600, 700},
		},
	}
	t.Run("master music all fetch", func(t *testing.T) {
		gotMusics, err := musicRepo.Fetch(ctx)
		// Basically does not occur
		if err != nil {
			t.Errorf("MusicRepository.Fetch() error = %v", err)
			return
		}
		if !assert.Equal(t, gotMusics, want) {
			t.Errorf("MusicRepository.Fetch() = %v, want %v", gotMusics, want)
		}
	})
}
