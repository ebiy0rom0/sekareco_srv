package logic

import (
	"sekareco_srv/domain/model"
	"sekareco_srv/infra"

	"github.com/pkg/errors"
)

type RecordLogic struct {
	recordRepo model.RecordRepository
}

func NewRecordLogic(r model.RecordRepository) model.RecordLogic {
	return &RecordLogic{
		recordRepo: r,
	}
}

func (l *RecordLogic) Store(r model.Record) (recordID int, err error) {
	if err = l.recordRepo.StartTransaction(); err != nil {
		infra.Logger.Error(errors.Wrap(err, "failed to start transaction"))
		return
	}

	if recordID, err = l.recordRepo.Store(r); err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to regist record: %#v", r))
		l.recordRepo.Rollback()
		return
	}

	l.recordRepo.Commit()
	return
}

func (l *RecordLogic) Update(personID int, musicID int, r model.Record) (err error) {
	if err = l.recordRepo.StartTransaction(); err != nil {
		infra.Logger.Error(errors.Wrap(err, "failed to start transaction"))
		return
	}

	if err = l.recordRepo.Update(personID, musicID, r); err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to modify record: %#v", r))
		l.recordRepo.Rollback()
		return
	}

	l.recordRepo.Commit()
	return
}

func (l *RecordLogic) GetByPersonID(personID int) (records []model.Record, err error) {
	if records, err = l.recordRepo.GetByPersonID(personID); err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to select record: person_id=%d", personID))
	}
	return
}
