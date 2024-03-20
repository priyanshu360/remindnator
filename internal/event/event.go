package event

import (
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/internal/schedule"
)

type Event interface {
	String() string
	Done()
	IsDone() bool
	NextSchedule() time.Time
}

type event struct {
	name          string
	start         time.Time
	end           time.Time
	schedule      schedule.Schedule
	isDone        bool
	lastPublished time.Time
}

func New(n string, st time.Time, et time.Time, d bool) *event {
	year, month, day := time.Now().Date()
	return &event{
		name:          n,
		start:         st,
		end:           et,
		isDone:        d,
		lastPublished: time.Date(year, month, day, 0, 0, 0, 0, time.UTC),
	}
}

func (e *event) String() string {
	// TODO #12 : event string function change according to presence of start/end
	return fmt.Sprintf("%s -> Starts at %s ", e.name, e.start)
}

func (e *event) Done() {
	e.isDone = true
}

func (e *event) IsDone() bool {
	return e.isDone || e.end.Before(time.Now())
}

func (e *event) NextSchedule() time.Time {
	return e.schedule.NextSchedule(e.lastPublished)
}
