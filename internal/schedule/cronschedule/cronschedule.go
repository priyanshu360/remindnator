package cronschedule

import (
	"time"

	"github.com/gorhill/cronexpr"
)

type schedule struct {
	cron string
}

func New(cron string) *schedule {
	return &schedule{
		cron: cron,
	}
}

func (s schedule) NextSchedule(t time.Time) time.Time {
	return cronexpr.MustParse(s.cron).Next(t)
}
