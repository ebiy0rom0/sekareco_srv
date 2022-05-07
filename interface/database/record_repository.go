package database

import "sekareco_srv/domain"

type RecordRepository struct {
	Handler SqlHandler
}

func (repository *RecordRepository) Regist(r domain.Record) (recordId int, err error) {
	// TODO: wip
	result, err := repository.Handler.Execute("INSERT INTO record VALUES ", r)
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

func (repository *RecordRepository) SelectArray(personId int) (recordList domain.RecordList, err error) {
	rows, err := repository.Handler.Query("SELECT record_id, person_id, music_id, record_easy, record_normal, record_hard, record_expert, record_master FROM record WHERE person_id = ?", personId)
	if err != nil {
		return
	}
	rows.Close()

	for rows.Next() {
		var record domain.Record
		err = rows.Scan(&record.RecordId, &record.PersonId, &record.MusicId, &record.RecordEasy, &record.RecordNormal, &record.RecordHard, &record.RecordExpert, &record.RecordMaster)
		if err != nil {
			return
		}

		recordList = append(recordList, record)
	}

	return
}
