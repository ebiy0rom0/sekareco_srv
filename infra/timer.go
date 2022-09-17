package infra

import (
	"sekareco_srv/domain/infra"
	"time"
)

type timeManager struct {
	timer *time.Location
}

// timer initialize.
// Set location(JST) and make Timer instance.
func init() {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	infra.Timer = &timeManager{
		timer: jst,
	}
}

// NowTime returns corrent time at JST.
func (t *timeManager) NowTime() time.Time {
	return time.Now().In(t.timer)
}

// NowDatetime returns current time and time in 'YYYY-MM-DD HH:ii:ss' format.
func (t *timeManager) NowDatetime() string {
	return t.NowTime().Format("2006-01-02 15:04:05")
}

// NowTimestamp returns current time and time in Unix-Timestamp.
func (t *timeManager) NowTimestamp() int64 {
	return t.NowTime().Unix()
}

// Add returns the time current + d.
func (t *timeManager) Add(d time.Duration) time.Time {
	return t.NowTime().Add(d)
}

// Sub returns the duration current - u.
func (t *timeManager) Sub(u time.Time) time.Duration {
	return t.NowTime().Sub(u)
}

// Before reports whether the current time is before u.
func (t *timeManager) Before(u time.Time) bool {
	return t.NowTime().Before(u)
}

var _ infra.ITimer = (*timeManager)(nil)
