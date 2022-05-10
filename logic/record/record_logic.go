package logic

import "sekareco_srv/domain"

type RecordLogic struct {
	Repository RecordRepository
}

func (logic *RecordLogic) RegistRecord(r domain.Record) (recordId int, err error) {
	recordId, err = logic.Repository.RegistRecord(r)
	return
}

func (logic *RecordLogic) ModifyRecord(personId int, musicId int, r domain.Record) (err error) {
	err = logic.Repository.ModifyRecord(r)
	return
}

func (logic *RecordLogic) GetPersonRecordList(personId int) (recordList domain.RecordList, err error) {
	recordList, err = logic.Repository.GetPersonRecordList(personId)
	return
}
