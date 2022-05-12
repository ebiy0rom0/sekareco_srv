package record

import "sekareco_srv/domain/model"

type RecordLogic struct {
	Repository RecordRepository
}

func (logic *RecordLogic) RegistRecord(r model.Record) (recordId int, err error) {
	recordId, err = logic.Repository.RegistRecord(r)
	return
}

func (logic *RecordLogic) ModifyRecord(personId int, musicId int, r model.Record) (err error) {
	err = logic.Repository.ModifyRecord(r)
	return
}

func (logic *RecordLogic) GetPersonRecordList(personId int) (recordList model.RecordList, err error) {
	recordList, err = logic.Repository.GetPersonRecordList(personId)
	return
}
