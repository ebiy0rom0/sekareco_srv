package logic

import "sekareco_srv/domain"

type RecordRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	RegistRecord(domain.Record) (int, error)
	ModifyRecord(domain.Record) error
	GetPersonRecordList(int) (domain.RecordList, error)
}
