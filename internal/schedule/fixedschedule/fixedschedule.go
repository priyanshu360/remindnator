package fixedschedule

import "time"

type schedule struct {
	next time.Time
}

func New(next time.Time) *schedule {
	return &schedule{
		next: next,
	}
}

func (s schedule) NextSchedule(time.Time) time.Time {
	return s.next
}
