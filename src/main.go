package main

import (
	"log"

	"github.com/priyanshu360/remindnator/src/sink/slackmessage"
	gcal "github.com/priyanshu360/remindnator/src/source/googlecalendar"
	"github.com/priyanshu360/remindnator/src/source/googletask"
	"github.com/priyanshu360/remindnator/src/watcher"
)

func init() {
	oauth()
	slackmessage.Init()
	if err := gcal.Init(); err != nil {
		log.Fatal(err)
	}
	if err := googletask.Init(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	cal, err := gcal.New("priyanshurajput360@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	task, err := googletask.New("My Tasks")
	if err != nil {
		log.Fatal(err)
	}

	notifier := slackmessage.New("C04KQEF85D5", "*/5 * * * *")
	cal.Subscribe(notifier)
	task.Subscribe(notifier)

	w := watcher.NewWatcher()
	w.Subscribe(cal, task)

	w.Run()
}
