package record

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"github.com/pkg/errors"
)

type RecordLogic struct {
	Repository RecordRepository
}

func (logic *RecordLogic) RegistRecord(r model.Record) (recordID int, err error) {
	if recordID, err = logic.Repository.RegistRecord(r); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist record: %#v", r))
	}
	return
}

func (logic *RecordLogic) ModifyRecord(personID int, musicID int, r model.Record) (err error) {
	if err = logic.Repository.ModifyRecord(r); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to modify record: %#v", r))
	}
	return
}

func (logic *RecordLogic) GetPersonRecordList(personID int) (recordList model.RecordList, err error) {
	if recordList, err = logic.Repository.GetPersonRecordList(personID); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select record: person_id=%d", personID))
	}
	return
}
