package infra

type Timer interface {
	NowDatetime() string
	NowTimestamp() int64
}
