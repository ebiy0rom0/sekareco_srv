package database

import "sekareco_srv/domain/model"

type MusicRepository struct {
	Handler SqlHandler
}

func (repository *MusicRepository) SelectAll() (musicList model.MusicList, err error) {
	rows, err := repository.Handler.Query("SELECT music_id, artist_id, music_name, jacket_url, level_easy, level_normal, level_hard, level_expert, level_master FROM master_music")
	if err != nil {
		return
	}
	rows.Close()

	for rows.Next() {
		var music model.Music
		err = rows.Scan(&music.MusicId, &music.MusicName, &music.MusicName, &music.JacketUrl, &music.LevelEasy, &music.LevelNormal, &music.LevelHard, &music.LevelExpert, &music.LevelMaster)
		if err != nil {
			return
		}

		musicList = append(musicList, music)
	}
	return
}
