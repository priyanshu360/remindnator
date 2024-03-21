package source

import "github.com/priyanshu360/remindnator/internal/sink"

//go:generate go-enum --marshal
// ENUM(googleTask, googleCalendar, SlackChannel)
type SourceEnum int

type Source interface {
	String() string
	Fetch() error
	FetchAll() error
	Subscribe(...sink.Sink)
}
