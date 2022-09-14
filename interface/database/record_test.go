package database_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"
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
				ScoreEasy:    0,
				RecordNormal: model.RECORD_CLEAR,
				ScoreNormal:  500,
				RecordHard:   model.RECORD_FULL_COMBO,
				ScoreHard:    800,
				RecordExpert: model.RECORD_ALL_PERFECT,
				ScoreExpert:  1200,
				RecordMaster: model.RECORD_ALL_PERFECT,
				ScoreMaster:  1500,
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
				ScoreEasy:    850,
				RecordNormal: model.RECORD_CLEAR,
				ScoreNormal:  1150,
				RecordHard:   model.RECORD_CLEAR,
				ScoreHard:    1420,
				RecordExpert: model.RECORD_FULL_COMBO,
				ScoreExpert:  1790,
				RecordMaster: model.RECORD_ALL_PERFECT,
				ScoreMaster:  2100,
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
				ScoreEasy:    0,
				RecordNormal: model.RECORD_NO_PLAY,
				ScoreNormal:  0,
				RecordHard:   model.RECORD_NO_PLAY,
				ScoreHard:    0,
				RecordExpert: model.RECORD_NO_PLAY,
				ScoreExpert:  0,
				RecordMaster: model.RECORD_NO_PLAY,
				ScoreMaster:  0,
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
					ScoreEasy:    220,
					RecordNormal: model.RECORD_FULL_COMBO,
					ScoreNormal:  590,
					RecordHard:   model.RECORD_ALL_PERFECT,
					ScoreHard:    900,
					RecordExpert: model.RECORD_ALL_PERFECT,
					ScoreExpert:  1200,
					RecordMaster: model.RECORD_NO_PLAY,
					ScoreMaster:  0,
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
					ScoreEasy:    0,
					RecordNormal: model.RECORD_NO_PLAY,
					ScoreNormal:  0,
					RecordHard:   model.RECORD_NO_PLAY,
					ScoreHard:    0,
					RecordExpert: model.RECORD_NO_PLAY,
					ScoreExpert:  0,
					RecordMaster: model.RECORD_NO_PLAY,
					ScoreMaster:  0,
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
					ScoreEasy:    0,
					RecordNormal: model.RECORD_NO_PLAY,
					ScoreNormal:  0,
					RecordHard:   model.RECORD_NO_PLAY,
					ScoreHard:    0,
					RecordExpert: model.RECORD_NO_PLAY,
					ScoreExpert:  0,
					RecordMaster: model.RECORD_NO_PLAY,
					ScoreMaster:  0,
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
		wantRecords []outputdata.Record
		wantErr     bool
	}{
		{
			name: "select successfully at single record",
			args: args{
				personID: 1,
			},
			wantRecords: []outputdata.Record{
				{
					MusicID: 1,
					Records: []int{
						model.RECORD_CLEAR,
						model.RECORD_FULL_COMBO,
						model.RECORD_ALL_PERFECT,
						model.RECORD_ALL_PERFECT,
						model.RECORD_NO_PLAY,
					},
					Scores: []int{220, 590, 900, 1200, 0},
				}, // inserted after updated item in test
			},
			wantErr: false,
		},
		{
			name: "select successfully at multiple records",
			args: args{
				personID: 2,
			},
			wantRecords: []outputdata.Record{
				{
					MusicID: 1,
					Records: []int{
						model.RECORD_ALL_PERFECT,
						model.RECORD_ALL_PERFECT,
						model.RECORD_ALL_PERFECT,
						model.RECORD_FULL_COMBO,
						model.RECORD_FULL_COMBO,
					},
					Scores: []int{300, 600, 900, 1195, 1480},
				}, // item not operated
				{
					MusicID: 2,
					Records: []int{
						model.RECORD_NO_PLAY,
						model.RECORD_NO_PLAY,
						model.RECORD_NO_PLAY,
						model.RECORD_NO_PLAY,
						model.RECORD_NO_PLAY,
					},
					Scores: []int{0, 0, 0, 0, 0},
				}, // updated item in test
				{
					MusicID: 3,
					Records: []int{
						model.RECORD_CLEAR,
						model.RECORD_CLEAR,
						model.RECORD_CLEAR,
						model.RECORD_FULL_COMBO,
						model.RECORD_ALL_PERFECT,
					},
					Scores: []int{850, 1150, 1420, 1790, 2100},
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
