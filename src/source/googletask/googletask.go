package googletasks

import (
	"context"
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/event"
	"github.com/priyanshu360/remindnator/src/sink"

	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	gtasks "google.golang.org/api/tasks/v1"
)

type taskList struct {
	id            string
	name          string
	notifier      []sink.Sink
	tasks         []event.Event
	nextFetchTime time.Time
}

var tasksService *gtasks.Service

func Init() (err error) {
	tasksService, err = gtasks.NewService(context.Background(), option.WithHTTPClient(config.CLIENT))
	return err
}

func NewTaskList(id string) (*taskList, error) {
	tasks, err := tasksService.Tasks.List(id).Do()
	if err != nil {
		return nil, err
	}

	return &taskList{
		id:   id,
		name: tasks.Etag,
	}, nil
}

func (tl *taskList) String() string {
	return tl.name
}

func (tl *taskList) Fetch() error {
	today := time.Now().UTC().Format("2006-01-02") // Format: YYYY-MM-DD
	timeMin := today + "T00:00:00Z"
	timeMax := today + "T23:59:59Z"
	tasks, err := tasksService.Tasks.List(tl.id).Do(googleapi.QueryParameter("dueMax", timeMax), googleapi.QueryParameter("dueMin", timeMin))

	if err != nil {
		return err
	}

	tl.tasks = make([]event.Event, 0)
	for _, t := range tasks.Items {
		// Process each task item
		fmt.Println(t)
	}
	return nil
}

func (tl *taskList) FetchAll() error {
	return nil
}

func (tl *taskList) Publish() error {
	for _, notif := range tl.notifier {
		notif.Publish(tl.tasks)
	}
	return nil
}

func (tl *taskList) Subscribe(notifier sink.Sink) {
	tl.notifier = append(tl.notifier, notifier)
}
