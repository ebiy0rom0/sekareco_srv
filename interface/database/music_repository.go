package database

import (
	"database/sql"
	"sekareco_srv/domain/model"

	"github.com/pkg/errors"
)

type MusicRepository struct {
	SqlHandler
}

func NewMusicRepository(h SqlHandler) model.MusicRepository {
	return &MusicRepository{h}
}

func (r *MusicRepository) Fetch() (musics []model.Music, err error) {
	rows, err := r.Query("SELECT music_id, artist_id, music_name, jacket_url, level_easy, level_normal, level_hard, level_expert, level_master FROM master_music")
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
