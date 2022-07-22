package interactor

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"

	"github.com/pkg/errors"
)

type recordInteractor struct {
	record      database.RecordRepository
	transaction database.SqlTransaction
}

func NewRecordInteractor(r database.RecordRepository, tx database.SqlTransaction) inputport.RecordInputport {
	return &recordInteractor{
		record:      r,
		transaction: tx,
	}
}

func (l *recordInteractor) Store(ctx context.Context, r inputdata.PostRecord) (int, error) {
	v, err := l.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		recordID, err := l.record.Store(ctx, model.Record{
			RecordID:     r.RecordID,
			PersonID:     r.PersonID,
			MusicID:      r.MusicID,
			RecordEasy:   r.RecordEasy,
			RecordNormal: r.RecordNormal,
			RecordHard:   r.RecordHard,
			RecordExpert: r.RecordExpert,
			RecordMaster: r.RecordMaster,
		})
		if err != nil {
			infra.Logger.Error(errors.Wrapf(err, "failed to regist record: %#v", r))
		}
		return recordID, err
	})

	return v.(int), err
}

func (l *recordInteractor) Update(ctx context.Context, personID int, musicID int, r inputdata.PutRecord) error {
	_, err := l.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		err := l.record.Update(ctx, personID, musicID, model.Record{
			RecordID: r.RecordID,
			PersonID: r.PersonID,
		})
		if err != nil {
			infra.Logger.Error(errors.Wrapf(err, "failed to modify record: %#v", r))
		}

		return nil, err
	})
	return err
}

func (l *recordInteractor) GetByPersonID(ctx context.Context, personID int) (records []model.Record, err error) {
	if records, err = l.record.GetByPersonID(ctx, personID); err != nil {
		infra.Logger.Error(errors.Wrapf(err, "failed to select record: person_id=%d", personID))
	}
	return
}
