package database

import (
	"context"
	"database/sql"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"

	"github.com/pkg/errors"
)

type musicRepository struct {
	infra.SqlHandler
}

func NewMusicRepository(h infra.SqlHandler) database.MusicRepository {
	return &musicRepository{h}
}

func (r *musicRepository) Fetch(ctx context.Context) (musics []outputdata.Music, err error) {
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
			&music.ArtistID,
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

		//convert to response data struct
		ret := outputdata.Music{
			MusicID:   music.MusicID,
			ArtistID:  music.ArtistID,
			MusicName: music.MusicName,
			JacketURL: music.JacketURL,
		}
		ret.Level = append(ret.Level, music.LevelEasy, music.LevelNormal, music.LevelHard, music.LevelExpert, music.LevelMaster)
		musics = append(musics, ret)
	}

	if len(musics) == 0 {
		err = errors.Wrap(sql.ErrNoRows, "failed")
	}
	return
}

// interface implemention checks
var _ database.MusicRepository = &musicRepository{}
