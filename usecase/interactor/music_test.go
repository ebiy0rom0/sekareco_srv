package interactor_test

import (
	"context"
	"reflect"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"
	"testing"
)

var mi inputport.MusicInputport

// this test case is same to repository test case
func Test_musicInteractor_Fetch(t *testing.T) {
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
		gotMusics, err := mi.Fetch(ctx)
		// Basically does not occur
		if err != nil {
			t.Errorf("musicInteractor.Fetch() error = %v", err)
			return
		}
		if !reflect.DeepEqual(gotMusics, want) {
			t.Errorf("musicInteractor.Fetch() = %v, want %v", gotMusics, want)
		}
	})
}
