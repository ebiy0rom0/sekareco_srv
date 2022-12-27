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
	query := "SELECT "
	query += "  music_id,   "
	query += "  artist_id,  "
	query += "  music_name, "
	query += "  jacket_url, "
	query += "  level_easy,   notes_easy,   "
	query += "  level_normal, notes_normal, "
	query += "  level_hard,   notes_hard,   "
	query += "  level_expert, notes_expert, "
	query += "  level_master, notes_master  "
	query += "FROM "
	query += "  master_music;"
	rows, err := r.Query(ctx, query)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer rows.Close()

	var musics []outputdata.Music
	for rows.Next() {
		var music model.Music
		err = rows.Scan(
			&music.MusicID,
			&music.ArtistID,
			&music.MusicName,
			&music.JacketURL,
			&music.LevelEasy,
			&music.NotesEasy,
			&music.LevelNormal,
			&music.NotesNormal,
			&music.LevelHard,
			&music.NotesHard,
			&music.LevelExpert,
			&music.NotesExpert,
			&music.LevelMaster,
			&music.NotesMaster,
		)
		if err != nil {
			return nil, errors.New(err.Error())
		}

		//convert to response data struct
		ret := outputdata.Music{
			MusicID:   music.MusicID,
			ArtistID:  music.ArtistID,
			MusicName: music.MusicName,
			JacketURL: music.JacketURL,
		}
		ret.Level = append([]int{}, music.LevelEasy, music.LevelNormal, music.LevelHard, music.LevelExpert, music.LevelMaster)
		ret.Notes = append([]int{}, music.NotesEasy, music.NotesNormal, music.NotesHard, music.NotesExpert, music.NotesMaster)
		musics = append(musics, ret)
	}

	if len(musics) == 0 {
		return nil, errors.New(sql.ErrNoRows.Error())
	}
	return musics, nil
}

var _ database.MusicRepository = (*musicRepository)(nil)
