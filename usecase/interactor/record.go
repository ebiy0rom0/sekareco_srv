package interactor

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"

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

func (l *recordInteractor) Store(ctx context.Context, personID int, r inputdata.AddRecord) (model.Record, error) {
	record := model.Record{
		PersonID:     personID,
		MusicID:      r.MusicID,
		RecordEasy:   r.RecordEasy,
		ScoreEasy:    r.ScoreEasy,
		RecordNormal: r.RecordNormal,
		ScoreNormal:  r.ScoreNormal,
		RecordHard:   r.RecordHard,
		ScoreHard:    r.ScoreHard,
		RecordExpert: r.RecordExpert,
		ScoreExpert:  r.ScoreExpert,
		RecordMaster: r.RecordMaster,
		ScoreMaster:  r.ScoreMaster,

	}

	v, err := l.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		recordID, err := l.record.Store(ctx, record)
		if err != nil {
			err = errors.Wrapf(err, "failed to regist record: %#v", r)
		}
		return recordID, err
	})

	if err != nil {
		return model.Record{}, err
	}

	record.RecordID = v.(int)
	return record, nil
}

func (l *recordInteractor) Update(ctx context.Context, personID int, musicID int, r inputdata.UpdateRecord) error {
	_, err := l.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		err := l.record.Update(ctx, personID, musicID, model.Record{
			PersonID:     personID,
			MusicID:      musicID,
			RecordEasy:   r.RecordEasy,
			ScoreEasy:    r.ScoreEasy,
			RecordNormal: r.RecordNormal,
			ScoreNormal:  r.ScoreNormal,
			RecordHard:   r.RecordHard,
			ScoreHard:    r.ScoreHard,
			RecordExpert: r.RecordExpert,
			ScoreExpert:  r.ScoreExpert,
			RecordMaster: r.RecordMaster,
			ScoreMaster:  r.ScoreMaster,
		})
		if err != nil {
			err = errors.Wrapf(err, "failed to modify record: %#v", r)
		}

		return nil, err
	})
	return err
}

func (l *recordInteractor) GetByPersonID(ctx context.Context, personID int) (records []outputdata.Record, err error) {
	if records, err = l.record.GetByPersonID(ctx, personID); err != nil {
		err = errors.Wrapf(err, "failed to select record: person_id=%d", personID)
	}
	return
}

var _ inputport.RecordInputport = (*recordInteractor)(nil)
