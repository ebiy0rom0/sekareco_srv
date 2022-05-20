package database

import (
	"sekareco_srv/domain/model"
)

type RecordRepository struct {
	Handler SqlHandler
}

func (repository *RecordRepository) StartTransaction() error {
	return repository.Handler.StartTransaction()
}

func (repository *RecordRepository) Commit() error {
	return repository.Handler.Commit()
}

func (repository *RecordRepository) Rollback() error {
	return repository.Handler.Rollback()
}

func (repository *RecordRepository) RegistRecord(r model.Record) (recordId int, err error) {
	query := "INSERT INTO record (person_id, music_id, record_easy, record_nomarl, record_hard, record_expert, record_master)"
	query += " VALUES (?, ?, ?, ?, ?, ?, ?);"
	result, err := repository.Handler.Execute(query, r.PersonId, r.MusicId, r.RecordEasy, r.RecordNormal, r.RecordHard, r.RecordExpert, r.RecordMaster)
	if err != nil {
		return
	}

	newId64, err := result.LastInsertId()
	if err != nil {
		return
	}

	recordId = int(newId64)
	return
}

func (repository *RecordRepository) ModifyRecord(r model.Record) (err error) {
	query := "UPDATE record SET record_easy = ?, record_normal = ?, record_hard = ?, record_expert = ?, record_master = ? WHERE person_id = ? AND music_id = ?;"
	_, err = repository.Handler.Execute(query, r.RecordEasy, r.RecordNormal, r.RecordNormal, r.RecordExpert, r.RecordMaster, r.PersonId, r.MusicId)
	return
}

func (repository *RecordRepository) GetPersonRecordList(personId int) (recordList model.RecordList, err error) {
	query := "SELECT person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master FROM record WHERE person_id = ?;"
	rows, err := repository.Handler.Query(query, personId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Record
		err = rows.Scan(&record.PersonId, &record.MusicId, &record.RecordEasy, &record.RecordNormal, &record.RecordHard, &record.RecordExpert, &record.RecordMaster)
		if err != nil {
			return
		}

		recordList = append(recordList, record)
	}

	return
}
