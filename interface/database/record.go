package database

import (
	"context"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"
)

type recordRepository struct {
	infra.SqlHandler
}

func NewRecordRepository(h infra.SqlHandler) *recordRepository {
	return &recordRepository{h}
}

func (r *recordRepository) Store(ctx context.Context, rec model.Record) (recordID int, err error) {
	query := "INSERT INTO record ("
	query += "  person_id, "
	query += "  music_id, "
	query += "  record_easy,   score_easy,   "
	query += "  record_normal, score_normal, "
	query += "  record_hard,   score_hard,   "
	query += "  record_expert, score_expert, "
	query += "  record_master, score_master  "
	query += ") VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	result, err := dao.Execute(ctx, query,
		rec.PersonID,
		rec.MusicID,
		rec.RecordEasy,
		rec.ScoreEasy,
		rec.RecordNormal,
		rec.ScoreNormal,
		rec.RecordHard,
		rec.ScoreHard,
		rec.RecordExpert,
		rec.ScoreExpert,
		rec.RecordMaster,
		rec.ScoreMaster,
	)
	if err != nil {
		return
	}

	newID64, err := result.LastInsertId()
	if err != nil {
		return
	}

	recordID = int(newID64)
	return
}

func (r *recordRepository) Update(ctx context.Context, personID int, musicID int, rec model.Record) (err error) {
	query := "UPDATE record "
	query += "SET "
	query += "  record_easy   = ?, score_easy =   ?, "
	query += "  record_normal = ?, score_normal = ?, "
	query += "  record_hard   = ?, score_hard   = ?, "
	query += "  record_expert = ?, score_expert = ?, "
	query += "  record_master = ?, score_master = ?  "
	query += "WHERE "
	query += "  person_id = ? AND music_id = ?;"

	dao, ok := getTx(ctx)
	if !ok {
		dao = r
	}

	_, err = dao.Execute(ctx, query,
		rec.RecordEasy,
		rec.ScoreEasy,
		rec.RecordNormal,
		rec.ScoreNormal,
		rec.RecordHard,
		rec.ScoreHard,
		rec.RecordExpert,
		rec.ScoreExpert,
		rec.RecordMaster,
		rec.ScoreMaster,
		personID,
		musicID,
	)
	return
}

func (r *recordRepository) GetByPersonID(ctx context.Context, personID int) (records []outputdata.Record, err error) {
	query := "SELECT "
	query += "  music_id, "
	query += "  record_easy,   score_easy,   "
	query += "  record_normal, score_normal, "
	query += "  record_hard,   score_hard,   "
	query += "  record_expert, score_expert, "
	query += "  record_master, score_master  "
	query += "FROM "
	query += "  record "
	query += "WHERE "
	query += "  person_id = ?;"

	rows, err := r.Query(ctx, query, personID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var record model.Record
		err = rows.Scan(
			&record.MusicID,
			&record.RecordEasy,
			&record.ScoreEasy,
			&record.RecordNormal,
			&record.ScoreNormal,
			&record.RecordHard,
			&record.ScoreHard,
			&record.RecordExpert,
			&record.ScoreExpert,
			&record.RecordMaster,
			&record.ScoreMaster,
		)

		if err != nil {
			return
		}

		//convert to response data struct
		ret := outputdata.Record{
			MusicID: record.MusicID,
		}

		ret.Records = append([]int{}, record.RecordEasy, record.RecordNormal, record.RecordHard, record.RecordExpert, record.RecordMaster)
		ret.Scores = append([]int{}, record.ScoreEasy, record.ScoreNormal, record.ScoreHard, record.ScoreExpert, record.ScoreMaster)
		records = append(records, ret)
	}
	return
}

// interface implementation checks
var _ database.RecordRepository = (*recordRepository)(nil)
