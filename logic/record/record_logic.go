package record

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"
)

type RecordLogic struct {
	Repository RecordRepository
}

func (logic *RecordLogic) RegistRecord(r model.Record) (recordId int, err error) {
	if recordId, err = logic.Repository.RegistRecord(r); err != nil {
		logger.Logger.Error(err)
	}
	return
}

func (logic *RecordLogic) ModifyRecord(personId int, musicId int, r model.Record) (err error) {
	if err = logic.Repository.ModifyRecord(r); err != nil {
		logger.Logger.Error(err)
	}
	return
}

func (logic *RecordLogic) GetPersonRecordList(personId int) (recordList model.RecordList, err error) {
	if recordList, err = logic.Repository.GetPersonRecordList(personId); err != nil {
		logger.Logger.Error(err)
	}
	return
}
