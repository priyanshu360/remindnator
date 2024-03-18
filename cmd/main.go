package main

import (
	"fmt"
	"log"

	"github.com/priyanshu360/remindnator/internal/watcher"
	"github.com/priyanshu360/remindnator/pkg/sink/slackmessage"
	gcal "github.com/priyanshu360/remindnator/pkg/source/googlecalendar"
	"github.com/priyanshu360/remindnator/pkg/source/googletask"
	"github.com/priyanshu360/remindnator/util"
)

func init() {
	if err := util.LoadLocalEnvFile(); err != nil {
		fmt.Println(err)
	}
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
