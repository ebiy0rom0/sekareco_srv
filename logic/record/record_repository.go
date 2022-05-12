package record

import "sekareco_srv/domain/model"

type RecordRepository interface {
	StartTransaction() error
	Commit() error
	Rollback() error
	RegistRecord(model.Record) (int, error)
	ModifyRecord(model.Record) error
	GetPersonRecordList(int) (model.RecordList, error)
}
