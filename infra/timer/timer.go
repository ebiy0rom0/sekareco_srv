package timer

import (
	"sekareco_srv/domain/common"
	"time"
)

var Timer common.Timer

type TimeManager struct {
}

func InitTimer() {
	Timer = new(TimeManager)
}

func (_ *TimeManager) NowDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func (_ *TimeManager) NowTimestamp() time.Time {
	t := time.Now()
	return time.Unix(t.Unix(), 0)
}
