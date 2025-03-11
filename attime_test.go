package chronus

import (
	"testing"
	"time"
)

func TestAtTime(t *testing.T) {
	cron := New()
	now := time.Now()
	after := now.Add(time.Minute)
	cron.AddFunc("at "+after.Format("2006-01-02 15:04:05"), func() {
		t.Logf("since :%v", time.Since(now))
	})
	cron.Start()
	time.Sleep(time.Minute * 2)
}
