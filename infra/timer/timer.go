package timer

import (
	"sekareco_srv/domain/infra"
	"time"
)

var Timer infra.Timer

type TimeManager struct {
	timer *time.Location
}

func InitTimer() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	Timer = &TimeManager{
		timer: jst,
	}
}

func (t *TimeManager) NowDatetime() string {
	return time.Now().In(t.timer).Format("2006-01-02 15:04:05")
}

func (t *TimeManager) NowTimestamp() int64 {
	return time.Now().In(t.timer).Unix()
}
