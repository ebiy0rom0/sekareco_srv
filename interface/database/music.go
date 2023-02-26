package database

import (
	"context"
	"database/sql"
	"sekareco_srv/domain/model"
	"sekareco_srv/interface/infra"
	"sekareco_srv/usecase/database"
	"sekareco_srv/usecase/outputdata"

	"github.com/ebiy0rom0/errors"
)

type musicRepository struct {
	infra.SqlHandler
}

func NewMusicRepository(h infra.SqlHandler) *musicRepository {
	return &musicRepository{h}
}

func (r *musicRepository) Fetch(ctx context.Context) ([]outputdata.Music, error) {
	query := `SELECT * FROM master_music;`

	rows, err := r.QueryxContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()

	var music model.Music
	var musics []outputdata.Music
	for rows.Next() {
		if err := rows.StructScan(&music); err != nil {
			return nil, errors.WithStack(err)
		}
		//convert to response data struct
		ret := outputdata.Music{
			MusicID:   music.MusicID,
			GroupID:   music.GroupID,
			MusicName: music.MusicName,
			JacketURL: music.JacketURL,
		}
		ret.Level = append([]int{}, music.LevelEasy, music.LevelNormal, music.LevelHard, music.LevelExpert, music.LevelMaster)
		ret.Notes = append([]int{}, music.NotesEasy, music.NotesNormal, music.NotesHard, music.NotesExpert, music.NotesMaster)
		musics = append(musics, ret)
	}

	if len(musics) == 0 {
		return nil, errors.WithStack(sql.ErrNoRows)
	}
	return musics, nil
}

var _ database.MusicRepository = (*musicRepository)(nil)
