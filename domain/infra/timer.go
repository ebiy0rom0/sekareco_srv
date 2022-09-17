package infra

import "time"

type ITimer interface {
	NowTime() time.Time
	NowDatetime() string
	NowTimestamp() int64
	Add(time.Duration) time.Time
	Sub(time.Time) time.Duration
	Before(time.Time) bool
}

// Timer instance to be called from code.
var Timer ITimer
