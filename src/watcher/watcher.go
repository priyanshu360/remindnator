package watcher

import (
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/source"
)

type Watcher struct {
	sources []source.Source
}

func (w *Watcher) Subscribe(s source.Source) {
	w.sources = append(w.sources, s)
}

func NewWatcher() *Watcher {
	return &Watcher{}
}

func (w *Watcher) Run() {
	for {
		for _, source := range w.sources {
			if err := source.Fetch(); err != nil {
				fmt.Println(err)
				continue
			}
			source.Publish()
		}
		time.Sleep(time.Duration(config.SLEEP_TIME) * time.Minute)
	}
}
