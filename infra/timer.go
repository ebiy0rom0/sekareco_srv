package infra

import (
	"time"
)

var Timer *timeManager

type timeManager struct {
	timer *time.Location
}

func init() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	Timer = &timeManager{
		timer: jst,
	}
}

func (t *timeManager) NowTime() time.Time {
	return time.Now().In(t.timer)
}

func (t *timeManager) NowDatetime() string {
	return t.NowTime().Format("2006-01-02 15:04:05")
}

func (t *timeManager) NowTimestamp() int64 {
	return t.NowTime().Unix()
}

func (t *timeManager) Add(d time.Duration) time.Time {
	return t.NowTime().Add(d)
}

func (t *timeManager) Sub(u time.Time) time.Duration {
	return t.NowTime().Sub(u)
}

func (t *timeManager) Before(u time.Time) bool {
	return t.NowTime().Before(u)
}
