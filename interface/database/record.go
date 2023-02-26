package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"

	"github.com/ebiy0rom0/errors"
)

type recordRepository struct {
	infra.SqlHandler
}

func NewRecordRepository(h infra.SqlHandler) *recordRepository {
	return &recordRepository{h}
}

func (r *recordRepository) Store(ctx context.Context, rec model.Record) (int, error) {
	query := `
	INSERT INTO person_record (
		person_id,     music_id,
		record_easy,   score_easy,
		record_normal, score_normal,
		record_hard,   score_hard,
		record_expert, score_expert,
		record_master, score_master
	) VALUES (
		:person_id,     :music_id,
		:record_easy,   :score_easy,
		:record_normal, :score_normal,
		:record_hard,   :score_hard,
		:record_expert, :score_expert,
		:record_master, :score_master
	);`

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.ExecNamedContext(ctx, query, rec)
	if err != nil {
		return 0, errors.WithStack(err)
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return 0, errors.WithStack(err)
	}

	return int(newID64), nil
}

func (r *recordRepository) Update(ctx context.Context, personID int, musicID int, rec model.Record) error {
	query := `
	UPDATE person_record
	SET
	  record_easy   = :record_easy,   score_easy =   :score_easy,
	  record_normal = :record_normal, score_normal = :score_normal,
	  record_hard   = :record_hard,   score_hard   = :score_hard,
	  record_expert = :record_expert, score_expert = :score_expert,
	  record_master = :record_master, score_master = :score_master
	WHERE
	  person_id = :person_id AND
	  music_id = :music_id;`

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	if _, err := dao.ExecNamedContext(ctx, query, rec); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *recordRepository) GetByPersonID(ctx context.Context, personID int) ([]outputdata.Record, error) {
	query := `SELECT * FROM person_record WHERE person_id = $1;`

	rows, err := r.QueryxContext(ctx, query, personID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var record model.Record
	var records []outputdata.Record
	for rows.Next() {
		if err := rows.StructScan(&record); err != nil {
			return nil, errors.WithStack(err)
		}
		//convert to response data struct
		ret := outputdata.Record{
			MusicID: record.MusicID,
		}

		ret.Records = append([]int{}, record.RecordEasy, record.RecordNormal, record.RecordHard, record.RecordExpert, record.RecordMaster)
		ret.Scores = append([]int{}, record.ScoreEasy, record.ScoreNormal, record.ScoreHard, record.ScoreExpert, record.ScoreMaster)
		records = append(records, ret)
	}
	return records, nil
}

var _ database.RecordRepository = (*recordRepository)(nil)
