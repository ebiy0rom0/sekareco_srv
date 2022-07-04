package record

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra/logger"

	"github.com/pkg/errors"
)

type RecordLogic struct {
	Repository RecordRepository
}

func (l *RecordLogic) RegistRecord(r model.Record) (recordID int, err error) {
	if recordID, err = l.Repository.RegistRecord(r); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to regist record: %#v", r))
	}
	return
}

func (l *RecordLogic) ModifyRecord(personID int, musicID int, r model.Record) (err error) {
	if err = l.Repository.ModifyRecord(r); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to modify record: %#v", r))
	}
	return
}

func (l *RecordLogic) GetPersonRecordList(personID int) (recordList model.RecordList, err error) {
	if recordList, err = l.Repository.GetPersonRecordList(personID); err != nil {
		logger.Logger.Error(errors.Wrapf(err, "failed to select record: person_id=%d", personID))
	}
	return
}
