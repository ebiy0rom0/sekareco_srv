package logic

import "sekareco_srv/domain"

type MusicRepository interface {
	SelectAll() (domain.MusicList, error)
}
