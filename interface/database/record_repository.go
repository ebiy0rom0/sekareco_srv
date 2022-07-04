package database

import (
	"sekareco_srv/domain/model"
)

type RecordRepository struct {
	Handler SqlHandler
}

func (r *RecordRepository) StartTransaction() error {
	return r.Handler.StartTransaction()
}

func (r *RecordRepository) Commit() error {
	return r.Handler.Commit()
}

func (r *RecordRepository) Rollback() error {
	return r.Handler.Rollback()
}

func (r *RecordRepository) RegistRecord(rec model.Record) (recordID int, err error) {
	query := "INSERT INTO record (person_id, music_id, record_easy, record_nomarl, record_hard, record_expert, record_master)"
	query += " VALUES (?, ?, ?, ?, ?, ?, ?);"
	result, err := r.Handler.Execute(query,
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

func (r *RecordRepository) ModifyRecord(rec model.Record) (err error) {
	query := "UPDATE record SET record_easy = ?, record_normal = ?, record_hard = ?, record_expert = ?, record_master = ? WHERE person_id = ? AND music_id = ?;"
	_, err = r.Handler.Execute(query,
		rec.RecordEasy,
		rec.RecordNormal,
		rec.RecordNormal,
		rec.RecordExpert,
		rec.RecordMaster,
		rec.PersonID,
		rec.MusicID,
	)
	return
}

func (rec *RecordRepository) GetPersonRecordList(personID int) (recordList model.RecordList, err error) {
	query := "SELECT person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master FROM record WHERE person_id = ?;"
	rows, err := rec.Handler.Query(query, personID)
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

		recordList = append(recordList, record)
	}

	return
}
