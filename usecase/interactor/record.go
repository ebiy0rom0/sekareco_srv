package interactor

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/inputdata"
	"sekareco_srv/usecase/inputport"
	"sekareco_srv/usecase/outputdata"

	"github.com/ebiy0rom0/errors"
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
			return nil, errors.Wrapf(err, "failed to store record:%+v", record)
		}
		return recordID, nil
	})

	if err != nil {
		return model.Record{}, err
	}

	record.RecordID = v.(int)
	return record, nil
}

func (l *recordInteractor) Update(ctx context.Context, personID int, musicID int, r inputdata.UpdateRecord) error {
	record := model.Record{
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
	}

	_, err := l.transaction.Do(ctx, func(ctx context.Context) (interface{}, error) {
		if err := l.record.Update(ctx, personID, musicID, record); err != nil {
			return nil, errors.Wrapf(err, "failed to update record:%+v", record)
		}
		return nil, nil
	})
	return err
}

func (l *recordInteractor) GetByPersonID(ctx context.Context, personID int) ([]outputdata.Record, error) {
	records, err := l.record.GetByPersonID(ctx, personID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to select record: person_id=%d", personID)
	}
	return records, nil
}

var _ inputport.RecordInputport = (*recordInteractor)(nil)
