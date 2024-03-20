package googlecalendar

import (
	"context"
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/internal/event"
	"github.com/priyanshu360/remindnator/internal/sink"

	gc "google.golang.org/api/calendar/v3"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
)

type source struct {
	id     string
	name   string
	sinks  []sink.Sink
	events []event.Event
}

var gcService *gc.Service

func Init() (err error) {
	// ctx := context.Background()
	// gcService, err = gc.NewService(ctx, option.WithAPIKey(config.API_KEY), option.WithScopes(gc.CalendarEventsReadonlyScope))
	gcService, err = gc.NewService(context.Background(), option.WithHTTPClient(config.CLIENT))
	return err
}

func New(id string) (*source, error) {
	events, err := gcService.Events.List(id).Do()
	if err != nil {
		return nil, err
	}

	return &source{
		id:   id,
		name: events.Summary,
	}, nil
}

func (gcal *source) String() string {
	return gcal.name
}

func (gcal *source) Fetch() error {
	today := time.Now().UTC().Format("2006-01-02") // Format: YYYY-MM-DD
	timeMin := today + "T00:00:00Z"
	timeMax := today + "T23:59:59Z"
	events, err := gcService.Events.List(gcal.id).Do(googleapi.QueryParameter("timeMax", timeMax), googleapi.QueryParameter("timeMin", timeMin))

	// events, err := gcService.Events.List("primary").Do()
	if err != nil {
		return err
	}

	gcal.events = make([]event.Event, 0)
	for _, e := range events.Items {
		st := time.Now()
		et := time.Now()
		if e.Start != nil {
			st, err = time.Parse(time.RFC3339, e.Start.DateTime)
			if err != nil {
				fmt.Println(err)
				continue
			}
			et, err = time.Parse(time.RFC3339, e.End.DateTime)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
		gcal.events = append(gcal.events, event.New(e.Summary, st, et, time.Now().After(et)))
	}
	return nil
}

func (gcal *source) FetchAll() error {
	return nil
}

func (gcal *source) Publish() error {
	for _, s := range gcal.sinks {
		s.Publish(gcal.events)
	}
	return nil
}

func (gcal *source) Subscribe(sinks ...sink.Sink) {
	gcal.sinks = append(gcal.sinks, sinks...)
}
