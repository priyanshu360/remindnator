package source

import "github.com/priyanshu360/remindnator/src/sink"

type Source interface {
	String() string
	Fetch() error
	FetchAll() error
	Publish() error
	Subscribe(...sink.Sink)
}
