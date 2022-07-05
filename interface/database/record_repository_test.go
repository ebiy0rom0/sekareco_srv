package database

import (
	"reflect"
	"sekareco_srv/domain/model"
	_database "sekareco_srv/logic/database"
	"testing"
)

func Test_Store(t *testing.T) {
	type args struct {
		rec model.Record
	}
	tests := []struct {
		name         string
		r            _database.RecordRepository
		args         args
		wantRecordID int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordID, err := tt.r.Store(tt.args.rec)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRecordID != tt.wantRecordID {
				t.Errorf("RecordRepository.Store() = %v, want %v", gotRecordID, tt.wantRecordID)
			}
		})
	}
}

func Test_Update(t *testing.T) {
	type args struct {
		personID int
		musicID  int
		rec      model.Record
	}
	tests := []struct {
		name    string
		r       _database.RecordRepository
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Update(tt.args.personID, tt.args.musicID, tt.args.rec); (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_GetByPersonID(t *testing.T) {
	type args struct {
		personID int
	}
	tests := []struct {
		name        string
		r           _database.RecordRepository
		args        args
		wantRecords []model.Record
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := tt.r.GetByPersonID(tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.GetByPersonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("RecordRepository.GetByPersonID() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
