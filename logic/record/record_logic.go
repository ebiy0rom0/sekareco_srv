package logic

import "sekareco_srv/domain"

type RecordLogic struct {
	Repository RecordRepository
}

func (logic *RecordLogic) RegistRecord(r domain.Record) (recordId int, err error) {
	recordId, err = logic.Repository.Regist(r)
	return
}

func (logic *RecordLogic) GetRecordList(personId int) (recordList domain.RecordList, err error) {
	recordList, err = logic.Repository.SelectArray(personId)
	return
}
