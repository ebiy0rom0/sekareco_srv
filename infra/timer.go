package infra

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

func (t *TimeManager) NowTime() time.Time {
	return time.Now().In(t.timer)
}

func (t *TimeManager) NowDatetime() string {
	return t.NowTime().Format("2006-01-02 15:04:05")
}

func (t *TimeManager) NowTimestamp() int64 {
	return t.NowTime().Unix()
}

func (t *TimeManager) Add(d time.Duration) time.Time {
	return t.NowTime().Add(d)
}

func (t *TimeManager) Before(u time.Time) bool {
	return t.NowTime().Before(u)
}
