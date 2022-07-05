package database_test

import (
	"sekareco_srv/domain/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fetch(t *testing.T) {
	tests := []struct {
		name       string
		r          model.MusicRepository
		wantMusics []model.Music
		wantErr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMusics, err := tt.r.Fetch()
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