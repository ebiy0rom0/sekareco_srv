package infra

import (
	"testing"
	"time"
)

// Since they can't be executed at same time,
// the difference is evaluated as passing if it is within 100 miliseconds.
var TOLERANCE_RANGE = 100 * time.Millisecond

var timer *timeManager

func init () {
	time.Local = time.FixedZone("Asia/Tokyo", 9*60*60)
	jst, _ := time.LoadLocation("Local")
	timer = &timeManager{
		timer: jst,
	}
}

func Test_timeManager_NowTime(t *testing.T) {
	// ??
	t.Run("now time check", func(t *testing.T) {
		d := time.Since(timer.NowTime())
		if d > TOLERANCE_RANGE {
			t.Error("timeManager.NowTime()")
		}
	})
}

func Test_timeManager_NowDatetime(t *testing.T) {
	// ??
	jst, _ := time.LoadLocation("Local")
	t.Run("datetime check", func(t *testing.T) {
		if timer.NowDatetime() != time.Now().In(jst).Format("2006-01-02 15:04:05") {
			t.Logf("my timer: %s, time: %s", timer.NowDatetime(), time.Now().Format("2006-01-02 15:04:05"))
			t.Error("timeManager.NowDatetime()")
		}
	})
}

func Test_timeManager_NowTimestamp(t *testing.T) {
	// ??
	t.Run("timestamp check", func(t *testing.T) {
		if timer.NowTimestamp() != time.Now().Unix() {
			t.Error("timeManager.NowTimestamp()")
		}
	})
}

func Test_timeManager_Add(t *testing.T) {
	// ??
	duration := 9 * time.Hour
	t.Run("time add check", func(t *testing.T) {
		utc := timer.Add(-duration)
		if time.Now().UTC().Equal(utc) {
			t.Error("timeManager.Add()")
		}
	})
}

func Test_timeManager_Sub(t *testing.T) {
	// ??
	tests := []struct{
		name     string
		duration time.Duration
	}{
		{
			name: "duration plus",
			duration: 5 * time.Hour,
		},
		{
			name: "duration minus",
			duration: -5 * time.Hour,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if d := timer.Sub(time.Now().Add(tt.duration)); d > -(tt.duration - TOLERANCE_RANGE) {
				t.Logf("failed duration: %d", timer.Sub(time.Now().Add(tt.duration)))
				t.Error("timeManager.Sub()")
			}
		})
	}
}

func Test_timeManager_Before(t *testing.T) {
	// ??
	t.Run("time before check", func(t *testing.T) {
		if timer.Before(time.Now().UTC()) {
			t.Errorf("timeManager.Before()")
		}
	})
}
