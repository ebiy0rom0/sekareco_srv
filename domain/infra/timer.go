package infra

import "time"

type Timer interface {
	NowTime() time.Time
	NowDatetime() string
	NowTimestamp() int64
	Add(time.Duration) time.Time
	Before(time.Time) bool
}
