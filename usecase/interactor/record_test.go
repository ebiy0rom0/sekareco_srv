package interactor_test

import (
	"context"
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"
	"testing"
)

var ri inputport.RecordInputport

func TestRecordInteractor_Store(t *testing.T) {
	ctx := context.Background()
	type args struct {
		personID int
		r        inputdata.AddRecord
	}
	tests := []struct {
		name    string
		args    args
		want    model.Record
		wantErr bool
	}{
		{
			name: "insert successfully",
			args: args{
				personID: 3,
				r: inputdata.AddRecord{
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
			},
			want: model.Record{
				RecordID:     5,
				PersonID:     3,
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
			wantErr: false,
		},
		{
			name: "[another person]insert successfully",
			args: args{
				personID: 3,
				r: inputdata.AddRecord{
					MusicID:      2,
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
			},
			want: model.Record{
				RecordID:     6,
				PersonID:     3,
				MusicID:      2,
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
			wantErr: false,
		},
		{
			name: "duplicate unique index",
			args: args{
				personID: 1,
				r: inputdata.AddRecord{
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
			},
			want:    model.Record{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ri.Store(ctx, tt.args.personID, tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recordInteractor.Store() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordInteractor_Update(t *testing.T) {
	ctx := context.Background()
	type args struct {
		personID int
		musicID  int
		r        inputdata.UpdateRecord
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
				r: inputdata.UpdateRecord{
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
				r: inputdata.UpdateRecord{
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
				r: inputdata.UpdateRecord{
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
			if err := ri.Update(ctx, tt.args.personID, tt.args.musicID, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRecordInteractor_GetByPersonID(t *testing.T) {
	ctx := context.Background()
	tests := []struct {
		name        string
		personID    int
		wantRecords []outputdata.Record
		wantErr     bool
	}{
		{
			name:     "select successfully at single record",
			personID: 1,
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
			name:     "select successfully at multiple records",
			personID: 2,
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecords, err := ri.GetByPersonID(ctx, tt.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("recordInteractor.GetByPersonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecords, tt.wantRecords) {
				t.Errorf("recordInteractor.GetByPersonID() = %v, want %v", gotRecords, tt.wantRecords)
			}
		})
	}
}
