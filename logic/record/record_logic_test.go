package record_test

import (
	"reflect"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/sql"
	"sekareco_srv/interface/database"
	"sekareco_srv/logic/record"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_RegistRecord(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql-mock: %s", err)
		return
	}

	l := &record.RecordLogic{
		Repository: &database.RecordRepository{
			Handler: &sql.SqlHandler{
				Conn: db,
				Tx:   nil,
			},
		},
	}

	r := model.Record{}
	if _, err := l.RegistRecord(r); err != nil {
		t.Errorf("regist record failed: %s", err)
	}

	type args struct {
		r model.Record
	}
	tests := []struct {
		name         string
		logic        *record.RecordLogic
		args         args
		wantRecordID int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordID, err := tt.logic.RegistRecord(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordLogic.RegistRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRecordID != tt.wantRecordID {
				t.Errorf("RecordLogic.RegistRecord() = %v, want %v", gotRecordID, tt.wantRecordID)
			}
		})
	}
}

func Test_ModifyRecord(t *testing.T) {
	type args struct {
		personID int
		musicID  int
		r        model.Record
	}
	tests := []struct {
		name    string
		logic   *record.RecordLogic
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.logic.ModifyRecord(tt.args.personID, tt.args.musicID, tt.args.r); (err != nil) != tt.wantErr {
				t.Errorf("RecordLogic.ModifyRecord() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_GetPersonRecordList(t *testing.T) {
	type args struct {
		personID int
	}
	tests := []struct {
		name           string
		logic          *record.RecordLogic
		args           args
		wantRecordList model.RecordList
		wantErr        bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordList, err := tt.logic.GetPersonRecordList(tt.args.personID)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordLogic.GetPersonRecordList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRecordList, tt.wantRecordList) {
				t.Errorf("RecordLogic.GetPersonRecordList() = %v, want %v", gotRecordList, tt.wantRecordList)
			}
		})
	}
}
