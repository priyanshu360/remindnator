package event

import (
	"fmt"
	"time"
)

type Event interface {
	String() string
	Done()
	IsDone() bool
}

type event struct {
	name   string
	time   time.Time
	isDone bool
}

func New(n string, t time.Time, d bool) *event {
	return &event{
		name:   n,
		time:   t,
		isDone: d,
	}
}

func (e *event) String() string {
	return fmt.Sprintf("%s -> Starts at %s ", e.name, e.time)
}

func (e *event) Done() {
	e.isDone = true
}

func (e *event) IsDone() bool {
	return e.isDone
}
