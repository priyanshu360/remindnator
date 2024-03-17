package sink

import "github.com/priyanshu360/remindnator/internal/event"

type Sink interface {
	// publish
	Publish([]event.Event)
}
