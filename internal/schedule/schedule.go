package schedule

import "time"

type Schedule interface {
	NextSchedule(time.Time) time.Time
}
