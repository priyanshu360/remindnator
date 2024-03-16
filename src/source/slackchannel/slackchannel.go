package slackchannel

import (
	"fmt"
	"time"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/event"
	"github.com/priyanshu360/remindnator/src/sink"
	"github.com/slack-go/slack"
)

type SlackChannel struct {
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

func New(channelId string) (*SlackChannel, error) {
	channel, err := slackBot.GetConversationInfo(&slack.GetConversationInfoInput{
		ChannelID:         channelId,
		IncludeLocale:     false,
		IncludeNumMembers: false,
	})
	if err != nil {
		return nil, err
	}
	return &SlackChannel{
		id:   channelId,
		name: channel.Name,
	}, nil
}

func (sc *SlackChannel) String() string {
	return sc.name
}

func (sc *SlackChannel) Fetch() error {
	today := time.Now().UTC().Format("2006-01-02")
	timeMin := today + "T00:00:00Z"
	timeMax := today + "T23:59:59Z"

	params := slack.GetConversationRepliesParameters{
		ChannelID: sc.id,
		Latest:    timeMax,
		Oldest:    timeMin,
	}

	replies, _, _, err := slackBot.GetConversationReplies(&params)
	if err != nil {
		return err
	}

	// Process fetched replies
	for _, reply := range replies {
		fmt.Println(reply.Text)
		// Process each reply as needed
	}

	return nil
}

func (sc *SlackChannel) FetchAll() error {
	// Implement fetching all events from Slack channel
	return nil
}

func (sc *SlackChannel) Publish() error {
	for _, sink := range sc.sinks {
		sink.Publish(sc.events)
	}
	return nil
}

func (sc *SlackChannel) Subscribe(sinks ...sink.Sink) {
	sc.sinks = append(sc.sinks, sinks...)
}
