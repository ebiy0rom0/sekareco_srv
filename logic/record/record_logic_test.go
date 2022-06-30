package record

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/sql"
	"sekareco_srv/interface/database"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_RegistRecord(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create sql-mock: %s", err)
		return
	}

	l := RecordLogic{
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
}
