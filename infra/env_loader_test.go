package infra

import "testing"

func TestLoadEnv(t *testing.T) {
	tests := []struct {
		name    string
		target  string
		wantErr bool
	}{
		{name: "load failure", target: ".env.nothing", wantErr: true},
		{name: "load successfully", target: ".env.testing", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadEnv(tt.target); (err != nil) != tt.wantErr {
				t.Errorf("LoadEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
