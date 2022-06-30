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

	mock.ExpectBegin()
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
	mock.ExpectCommit()

	repo.StartTransaction()
	if _, err := repo.RegistRecord(r); err != nil {
		t.Errorf("error: %s", err)
	}
	repo.Commit()
}
