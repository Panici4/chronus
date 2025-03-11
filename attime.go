package chronus

import "time"

type AtTimeSchedule struct {
	at time.Time
}

func AtTime(at time.Time) AtTimeSchedule {
	return AtTimeSchedule{
		at: at,
	}
}

// Next returns the next time this should be run.
func (schedule AtTimeSchedule) Next(t time.Time) time.Time {
	return schedule.at
}
