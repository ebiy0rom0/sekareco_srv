package common

import "time"

type Timer interface {
	NowDatetime() string
	NowTimestamp() time.Time
}
