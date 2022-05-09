package logic

import "sekareco_srv/domain"

type RecordRepository interface {
	Regist(domain.Record) (int, error)
	Modify(domain.Record) error
	SelectArray(int) (domain.RecordList, error)
}
