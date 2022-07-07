package database_test

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMusicRepository_Fetch(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name       string
		r          database.MusicRepository
		wantMusics []model.Music
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMusics, err := tt.r.Fetch(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("MusicRepository.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !assert.Equal(t, gotMusics, tt.wantMusics) {
				t.Errorf("MusicRepository.Fetch() = %v, want %v", gotMusics, tt.wantMusics)
			}
		})
	}
}
