package database_test

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

// initialize at TestMain() in login_test.go
var musicRepo database.MusicRepository

func TestMusicRepository_Fetch(t *testing.T) {
	ctx := context.Background()
	want := []model.Music{
		{
			MusicID:     1,
			ArtistID:    1,
			MusicName:   "test_music001",
			JacketURL:   "jacket/m_001.png",
			LevelEasy:   1,
			LevelNormal: 2,
			LevelHard:   3,
			LevelExpert: 4,
			LevelMaster: 5,
		},
		{
			MusicID:     2,
			ArtistID:    2,
			MusicName:   "test_music002",
			JacketURL:   "jacket/m_002.png",
			LevelEasy:   2,
			LevelNormal: 3,
			LevelHard:   4,
			LevelExpert: 5,
			LevelMaster: 6,
		},
		{
			MusicID:     3,
			ArtistID:    1,
			MusicName:   "test_music003",
			JacketURL:   "jacket/m_003.png",
			LevelEasy:   3,
			LevelNormal: 4,
			LevelHard:   5,
			LevelExpert: 6,
			LevelMaster: 7,
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
