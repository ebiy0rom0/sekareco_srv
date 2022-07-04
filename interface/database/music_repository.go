package database

import (
	"database/sql"
	"sekareco_srv/domain/model"

	"github.com/pkg/errors"
)

type MusicRepository struct {
	Handler SqlHandler
}

func (r *MusicRepository) SelectAll() (musicList model.MusicList, err error) {
	rows, err := r.Handler.Query("SELECT music_id, artist_id, music_name, jacket_url, level_easy, level_normal, level_hard, level_expert, level_master FROM master_music")
	if err != nil {
		err = errors.Wrap(err, "failed")
		return
	}
	defer rows.Close()

	for rows.Next() {
		var music model.Music
		err = rows.Scan(&music.MusicID, &music.MusicName, &music.MusicName, &music.JacketUrl, &music.LevelEasy, &music.LevelNormal, &music.LevelHard, &music.LevelExpert, &music.LevelMaster)
		if err != nil {
			err = errors.Wrap(err, "failed")
			return
		}

		musicList = append(musicList, music)
	}

	if len(musicList) == 0 {
		err = errors.Wrap(sql.ErrNoRows, "failed")
	}
	return
}
