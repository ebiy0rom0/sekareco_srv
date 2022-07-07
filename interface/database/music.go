package database

import (
	"context"
	"database/sql"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"

	"github.com/pkg/errors"
)

type musicRepository struct {
	infra.SqlHandler
}

func NewMusicRepository(h infra.SqlHandler) database.MusicRepository {
	return &musicRepository{h}
}

func (r *musicRepository) Fetch(ctx context.Context) (musics []model.Music, err error) {
	query := "SELECT music_id, artist_id, music_name, jacket_url, level_easy, level_normal, level_hard, level_expert, level_master FROM master_music"
	rows, err := r.Query(ctx, query)
	if err != nil {
		err = errors.Wrap(err, "failed")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var music model.Music
		err = rows.Scan(
			&music.MusicID,
			&music.MusicName,
			&music.MusicName,
			&music.JacketURL,
			&music.LevelEasy,
			&music.LevelNormal,
			&music.LevelHard,
			&music.LevelExpert,
			&music.LevelMaster,
		)
		if err != nil {
			err = errors.Wrap(err, "failed")
			return
		}

		musics = append(musics, music)
	}

	if len(musics) == 0 {
		err = errors.Wrap(sql.ErrNoRows, "failed")
	}
	return
}
