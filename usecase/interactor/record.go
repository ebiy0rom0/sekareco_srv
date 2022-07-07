package interactor

import (
	"context"
	"sekareco_srv/domain/model"
	_infra "sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputport"

	"github.com/pkg/errors"
)

type RecordInteractor struct {
	recordRepo  database.RecordRepository
	transaction database.SqlTransaction
}

func NewRecordInteractor(
	r database.RecordRepository,
	tx database.SqlTransaction,
) inputport.RecordInputport {
	return &RecordInteractor{
		recordRepo:  r,
		transaction: tx,
	}
}

func (l *RecordInteractor) Store(ctx context.Context, r model.Record) (recordID int, err error) {
	// v, err := l.transaction.Do(ctx, l.recordRepo.Store)
	// if err = l.recordRepo.StartTransaction(); err != nil {
	// 	_infra.Logger.Error(errors.Wrap(err, "failed to start transaction"))
	// 	return
	// }

	if recordID, err = l.recordRepo.Store(r); err != nil {
		_infra.Logger.Error(errors.Wrapf(err, "failed to regist record: %#v", r))
		// l.recordRepo.Rollback()
		return
	}

	// l.recordRepo.Commit()
	return
}

func (l *RecordInteractor) Update(personID int, musicID int, r model.Record) (err error) {
	// if err = l.recordRepo.StartTransaction(); err != nil {
	// 	_infra.Logger.Error(errors.Wrap(err, "failed to start transaction"))
	// 	return
	// }

	if err = l.recordRepo.Update(personID, musicID, r); err != nil {
		_infra.Logger.Error(errors.Wrapf(err, "failed to modify record: %#v", r))
		// l.recordRepo.Rollback()
		return
	}

	// l.recordRepo.Commit()
	return
}

func (l *RecordInteractor) GetByPersonID(personID int) (records []model.Record, err error) {
	if records, err = l.recordRepo.GetByPersonID(personID); err != nil {
		_infra.Logger.Error(errors.Wrapf(err, "failed to select record: person_id=%d", personID))
	}
	return
}
