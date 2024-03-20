package watcher

import (
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/internal/source"
)

type Watcher struct {
	sources []source.Source
}

func (w *Watcher) Subscribe(sources ...source.Source) {
	w.sources = append(w.sources, sources...)
}

func NewWatcher() *Watcher {
	return &Watcher{}
}

func (w *Watcher) Run() {
	for {
		for _, source := range w.sources {
			fmt.Println(source.String())
			if err := source.Fetch(); err != nil {
				fmt.Println(err)
				continue
			}
			source.Publish()
		}
		time.Sleep(time.Duration(config.SLEEP_TIME) * time.Minute)
	}
}
