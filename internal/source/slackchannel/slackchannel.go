package slackchannel

import (
	"log"
	"time"

	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/internal/event"
	"github.com/priyanshu360/remindnator/internal/sink"
	"github.com/priyanshu360/remindnator/internal/source"
	"github.com/priyanshu360/remindnator/util"
	"github.com/slack-go/slack"
)

type SlackChannel struct {
	Type      source.SourceEnum
	Title     string
	ID        string
	SinkIDs   []string
	ChannelID string
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
		Title:     channel.Name,
		ID:        util.GenerateUUID(5),
		ChannelID: channel.ID,
	}, nil
}

func (sc *SlackChannel) String() string {
	return sc.Title
}

func (sc *SlackChannel) Fetch() error {
	// TODO #1 : Debug oldest latest not working
	// today := time.Now().UTC().Format("2006-01-02")
	// timeMin := today + "T00:00:00Z"
	// timeMax := today + "T23:59:59Z"

	params := slack.GetConversationHistoryParameters{
		ChannelID: sc.ChannelID,
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
		sc.events = append(sc.events, event.New(reply.Text, time.Now(), time.Now(), false))
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
