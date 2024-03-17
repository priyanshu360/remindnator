package slackchannel

import (
	"log"
	"time"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/event"
	"github.com/priyanshu360/remindnator/src/sink"
	"github.com/slack-go/slack"
)

type slackChannel struct {
	name   string
	id     string
	sinks  []sink.Sink
	events []event.Event
}

var slackBot *slack.Client

func Init() error {
	slackBot = slack.New(config.SLACK_TOKEN)
	return nil
}

func New(channelId string) (*slackChannel, error) {
	channel, err := slackBot.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID:         channelId,
		IncludeLocale:     false,
		IncludeNumMembers: false,
	})
	if err != nil {
		return nil, err
	}

	return &slackChannel{
		name:   channel.Name,
		id:     channel.ID,
		sinks:  []sink.Sink{},
		events: []event.Event{},
	}, nil
}

func (sc *slackChannel) String() string {
	return sc.name
}

func (sc *slackChannel) Fetch() error {
	// TODO #1 : Debug oldest latest not working
	// today := time.Now().UTC().Format("2006-01-02")
	// timeMin := today + "T00:00:00Z"
	// timeMax := today + "T23:59:59Z"

	params := slack.GetConversationHistoryParameters{
		ChannelID: sc.id,
		// Latest:    timeMax,
		// Oldest: timeMin,
	}

	resp, err := slackBot.GetConversationHistory(&params)
	if err != nil {
		return err
	}

	log.Println(resp.Messages)
	// Process fetched replies
	for _, reply := range resp.Messages {
		// TODO #2 : parse and process
		sc.events = append(sc.events, event.NewEvent(reply.Text, time.Now(), false))
	}

	return nil
}

func (sc *slackChannel) FetchAll() error {
	// Implement fetching all events from Slack channel
	return nil
}

func (sc *slackChannel) Publish() error {
	for _, sink := range sc.sinks {
		sink.Publish(sc.events)
	}
	return nil
}

func (sc *slackChannel) Subscribe(sinks ...sink.Sink) {
	sc.sinks = append(sc.sinks, sinks...)
}
