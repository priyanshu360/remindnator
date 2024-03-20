package source

import "github.com/priyanshu360/remindnator/internal/sink"

type Source interface {
	String() string
	Fetch() error
	FetchAll() error
	Subscribe(...sink.Sink)
}
