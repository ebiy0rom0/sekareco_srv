package database

import (
	"sekareco_srv/domain/model"
)

type RecordRepository struct {
	SqlHandler
}

func NewRecordRepository(h SqlHandler) model.RecordRepository {
	return &RecordRepository{h}
}

func (r *RecordRepository) Store(rec model.Record) (recordID int, err error) {
	query := "INSERT INTO record (person_id, music_id, record_easy, record_nomarl, record_hard, record_expert, record_master)"
	query += " VALUES (?, ?, ?, ?, ?, ?, ?);"
	result, err := r.Execute(query,
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

func (r *RecordRepository) Update(personID int, musicID int, rec model.Record) (err error) {
	query := "UPDATE record SET record_easy = ?, record_normal = ?, record_hard = ?, record_expert = ?, record_master = ? WHERE person_id = ? AND music_id = ?;"
	_, err = r.Execute(query,
		rec.RecordEasy,
		rec.RecordNormal,
		rec.RecordNormal,
		rec.RecordExpert,
		rec.RecordMaster,
		personID,
		musicID,
	)
	return
}

func (r *RecordRepository) GetByPersonID(personID int) (records []model.Record, err error) {
	query := "SELECT person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master FROM record WHERE person_id = ?;"
	rows, err := r.Query(query, personID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Record
		err = rows.Scan(&record.PersonID, &record.MusicID, &record.RecordEasy, &record.RecordNormal, &record.RecordHard, &record.RecordExpert, &record.RecordMaster)
		if err != nil {
			return
		}

		records = append(records, record)
	}

	return
}
