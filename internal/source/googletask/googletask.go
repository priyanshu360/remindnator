package googletask

import (
	"context"
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/internal/event"
	"github.com/priyanshu360/remindnator/internal/sink"
	"github.com/priyanshu360/remindnator/internal/source"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	gtasks "google.golang.org/api/tasks/v1"
)

type TaskList struct {
	Type    source.SourceEnum
	ID      string
	TaskID  string
	Title   string
	SinkIDs []string
}

var tasksService *gtasks.Service

func Init() (err error) {
	tasksService, err = gtasks.NewService(context.Background(), option.WithHTTPClient(config.CLIENT))
	return err
}

func New(title string) (*taskList, error) {
	tasks, err := tasksService.Tasklists.List().Do()
	if err != nil {
		return nil, err
	}

	var id string
	for _, t := range tasks.Items {
		if title == t.Title {
			id = t.Id
			break
		}
	}

	if len(id) == 0 {
		return nil, fmt.Errorf("Invalid Task List Title %s", title)
	}

	return &taskList{
		id:    id,
		Title: title,
	}, nil
}

func (tl *taskList) String() string {
	return tl.Title
}

func (tl *taskList) Fetch() error {
	today := time.Now().UTC().Format("?2006-01-02") // Format: YYYY-MM-DD
	timeMin := today + "T00:00:00Z"
	timeMax := today + "T23:59:59Z"
	tasks, err := tasksService.Tasks.List(tl.id).Do(googleapi.QueryParameter("dueMax", timeMax), googleapi.QueryParameter("dueMin", timeMin))

	if err != nil {
		return err
	}

	tl.events = make([]event.Event, 0)
	for _, t := range tasks.Items {
		et, err := time.Parse(time.RFC3339, t.Due)
		if err != nil {
			fmt.Println(err)
			continue
		}
		st, err := time.Parse(time.RFC3339, t.Updated)
		if err != nil {
			fmt.Println(err)
			continue
		}
		tl.events = append(tl.events, event.New(t.Title, st, et, t.Completed != nil))
	}
	return nil
}

func (tl *taskList) FetchAll() error {
	return nil
}

func (tl *taskList) Publish() error {
	for _, notif := range tl.sinks {
		notif.Publish(tl.events)
	}
	return nil
}

func (tl *taskList) Subscribe(sinks ...sink.Sink) {
	tl.sinks = append(tl.sinks, sinks...)
}
