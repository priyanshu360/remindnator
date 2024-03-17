package sink

import "github.com/priyanshu360/remindnator/src/event"

type Sink interface {
	// publish
	Publish([]event.Event)
}
