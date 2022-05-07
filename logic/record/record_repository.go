package logic

import "sekareco_srv/domain"

type RecordRepository interface {
	Regist(domain.Record) (int, error)
	SelectArray(int) (domain.RecordList, error)
}
