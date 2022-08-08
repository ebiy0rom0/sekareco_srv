package database_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"testing"
)

var recordRepo database.RecordRepository

// this database check is unused transaction
func TestRecordRepository_Store(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name         string
		record       model.Record
		wantRecordID int
		wantErr      bool
	}{
		{
			name: "insert successfully",
			record: model.Record{
				PersonID:     1,
				MusicID:      1,
				RecordEasy:   model.RECORD_NO_PLAY,
				RecordNormal: model.RECORD_CLEAR,
				RecordHard:   model.RECORD_FULL_COMBO,
				RecordExpert: model.RECORD_ALL_PERFECT,
				RecordMaster: model.RECORD_ALL_PERFECT,
			},
			wantRecordID: 3,
			wantErr:      false,
		},
		{
			name: "[another person]insert successfully",
			record: model.Record{
				PersonID:     2,
				MusicID:      3,
				RecordEasy:   model.RECORD_CLEAR,
				RecordNormal: model.RECORD_CLEAR,
				RecordHard:   model.RECORD_CLEAR,
				RecordExpert: model.RECORD_FULL_COMBO,
				RecordMaster: model.RECORD_ALL_PERFECT,
			},
			wantRecordID: 4,
			wantErr:      false,
		},
		{
			name: "duplicate unique index",
			record: model.Record{
				PersonID:     1,
				MusicID:      1,
				RecordEasy:   model.RECORD_NO_PLAY,
				RecordNormal: model.RECORD_NO_PLAY,
				RecordHard:   model.RECORD_NO_PLAY,
				RecordExpert: model.RECORD_NO_PLAY,
				RecordMaster: model.RECORD_NO_PLAY,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordID, err := recordRepo.Store(ctx, tt.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && gotRecordID != tt.wantRecordID {
				t.Errorf("RecordRepository.Store() = %v, want %v", gotRecordID, tt.wantRecordID)
			}
		})
	}
}

// this database check is unused transaction
func TestRecordRepository_Update(t *testing.T) {
	ctx := context.Background()
	type args struct {
		personID int
		musicID  int
		rec      model.Record
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update successfully",
			args: args{
				personID: 1,
				musicID:  1,
				rec: model.Record{
					RecordEasy:   model.RECORD_CLEAR,
					RecordNormal: model.RECORD_FULL_COMBO,
					RecordHard:   model.RECORD_ALL_PERFECT,
					RecordExpert: model.RECORD_ALL_PERFECT,
					RecordMaster: model.RECORD_NO_PLAY,
				},
			},
			wantErr: false,
		},
		{
			name: "[another person]update successfully",
			args: args{
				personID: 2,
				musicID:  2,
				rec: model.Record{
					RecordEasy:   model.RECORD_NO_PLAY,
					RecordNormal: model.RECORD_NO_PLAY,
					RecordHard:   model.RECORD_NO_PLAY,
					RecordExpert: model.RECORD_NO_PLAY,
					RecordMaster: model.RECORD_NO_PLAY,
				},
			},
			wantErr: false,
		},
		{
			name: "un registered record update",
			args: args{
				personID: 1,
				musicID:  2,
				rec: model.Record{
					RecordEasy:   model.RECORD_NO_PLAY,
					RecordNormal: model.RECORD_NO_PLAY,
					RecordHard:   model.RECORD_NO_PLAY,
					RecordExpert: model.RECORD_NO_PLAY,
					RecordMaster: model.RECORD_NO_PLAY,
				},
			},
			wantErr: false,
		}, // not update, but not returned error because query is correct
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := recordRepo.Update(ctx, tt.args.personID, tt.args.musicID, tt.args.rec); (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecordRepository_GetByPersonID(t *testing.T) {
	ctx := context.Background()
	type args struct {
		personID int
	}
	tests := []struct {
		name        string
		args        args
		wantRecords []model.Record
		wantErr     bool
	}{
		{
			name: "select successfully at single record",
			args: args{
				personID: 1,
			},
			wantRecords: []model.Record{
				{
					PersonID:     1,
					MusicID:      1,
					RecordEasy:   model.RECORD_CLEAR,
					RecordNormal: model.RECORD_FULL_COMBO,
					RecordHard:   model.RECORD_ALL_PERFECT,
					RecordExpert: model.RECORD_ALL_PERFECT,
					RecordMaster: model.RECORD_NO_PLAY,
				}, // inserted after updated item in test
			},
			wantErr: false,
		},
		{
			name: "select successfully at multiple records",
			args: args{
				personID: 2,
			},
			wantRecords: []model.Record{
				{
					PersonID:     2,
					MusicID:      1,
					RecordEasy:   model.RECORD_ALL_PERFECT,
					RecordNormal: model.RECORD_ALL_PERFECT,
					RecordHard:   model.RECORD_ALL_PERFECT,
					RecordExpert: model.RECORD_FULL_COMBO,
					RecordMaster: model.RECORD_FULL_COMBO,
				}, // item not operated
				{
					PersonID:     2,
					MusicID:      2,
					RecordEasy:   model.RECORD_NO_PLAY,
					RecordNormal: model.RECORD_NO_PLAY,
					RecordHard:   model.RECORD_NO_PLAY,
					RecordExpert: model.RECORD_NO_PLAY,
					RecordMaster: model.RECORD_NO_PLAY,
				}, // updated item in test
				{
					PersonID:     2,
					MusicID:      3,
					RecordEasy:   model.RECORD_CLEAR,
					RecordNormal: model.RECORD_CLEAR,
					RecordHard:   model.RECORD_CLEAR,
					RecordExpert: model.RECORD_FULL_COMBO,
					RecordMaster: model.RECORD_ALL_PERFECT,
				}, // inserted item in test
			},
			wantErr: false,
		},
		{
			name: "un registered record person",
			args: args{
				personID: 3,
			},
			wantErr: false,
		}, // error not returned when select no rows in sqlHandler.Query()
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := recordRepo.GetByPersonID(ctx, tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.GetByPersonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("RecordRepository.GetByPersonID() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
