package database_test

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/sql"
	"sekareco_srv/interface/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_RegistRecord(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to new sql mock: %s", err)
	}
	defer db.Close()

	repo := &database.RecordRepository{
		Handler: &sql.SqlHandler{
			Conn: db,
			Tx:   nil,
		},
	}

	r := model.Record{
		PersonID:     1,
		MusicID:      1,
		RecordEasy:   model.RECORD_NO_PLAY,
		RecordNormal: model.RECORD_CLEAR,
		RecordHard:   model.RECORD_FULL_COMBO,
		RecordExpert: model.RECORD_ALL_PERFECT,
		RecordMaster: model.RECORD_ALL_PERFECT,
	}

	mock.ExpectPrepare("INSERT INTO record").
		ExpectExec().
		WithArgs(
			r.PersonID,
			r.MusicID,
			r.RecordEasy,
			r.RecordNormal,
			r.RecordHard,
			r.RecordExpert,
			r.RecordMaster,
		).
		WillReturnResult(sqlmock.NewResult(1, 1))

	repo.StartTransaction()
	if _, err := repo.RegistRecord(r); err != nil {
		t.Errorf("error: %s", err)
	}
	repo.Commit()
}

func TestRecordRepository_RegistRecord(t *testing.T) {
	type args struct {
		r model.Record
	}
	tests := []struct {
		name         string
		repository   *database.RecordRepository
		args         args
		wantRecordID int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRecordID, err := tt.repository.RegistRecord(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("RecordRepository.RegistRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRecordID != tt.wantRecordID {
				t.Errorf("RecordRepository.RegistRecord() = %v, want %v", gotRecordID, tt.wantRecordID)
			}
		})
	}
}
