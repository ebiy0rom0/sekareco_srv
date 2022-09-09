package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
)

type recordRepository struct {
	infra.SqlHandler
}

func NewRecordRepository(h infra.SqlHandler) *recordRepository {
	return &recordRepository{h}
}

func (r *recordRepository) Store(ctx context.Context, rec model.Record) (recordID int, err error) {
	query := "INSERT INTO record (person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master)"
	query += " VALUES (?, ?, ?, ?, ?, ?, ?);"

	dao, ok := GetTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.Execute(ctx, query,
		rec.PersonID,
		rec.MusicID,
		rec.RecordEasy,
		rec.RecordNormal,
		rec.RecordHard,
		rec.RecordExpert,
		rec.RecordMaster,
	)
	if err != nil {
		return
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return
	}

	recordID = int(newID64)
	return
}

func (r *recordRepository) Update(ctx context.Context, personID int, musicID int, rec model.Record) (err error) {
	query := "UPDATE record SET record_easy = ?, record_normal = ?, record_hard = ?, record_expert = ?, record_master = ? WHERE person_id = ? AND music_id = ?;"

	dao, ok := GetTx(ctx)
	if !ok {
		dao = r
	}

	_, err = dao.Execute(ctx, query,
		rec.RecordEasy,
		rec.RecordNormal,
		rec.RecordHard,
		rec.RecordExpert,
		rec.RecordMaster,
		personID,
		musicID,
	)
	return
}

func (r *recordRepository) GetByPersonID(ctx context.Context, personID int) (records []model.Record, err error) {
	query := "SELECT person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master FROM record WHERE person_id = ?;"
	rows, err := r.Query(ctx, query, personID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Record
		err = rows.Scan(
			&record.PersonID,
			&record.MusicID,
			&record.RecordEasy,
			&record.RecordNormal,
			&record.RecordHard,
			&record.RecordExpert,
			&record.RecordMaster,
		)

		if err != nil {
			return
		}
		records = append(records, record)
	}
	return
}

// interface implementation checks
var _ database.RecordRepository = (*recordRepository)(nil)
