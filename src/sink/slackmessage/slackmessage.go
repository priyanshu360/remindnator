package slackmessage

import (
	"log"

	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/event"

	"github.com/slack-go/slack"
)

type sink struct {
	channelId string
	crontab   string
}

var slackBot *slack.Client

func Init() error {
	slackBot = slack.New(config.SLACK_TOKEN)
	return nil
}

func New(id string, crontab string) *sink {
	return &sink{
		channelId: id,
		crontab:   crontab,
	}
}

func (s *sink) Publish(event []event.Event) {
	for _, e := range event {
		if err := sendSlackMessage(s.channelId, e.String()); err != nil {
			log.Fatal(err) // TODO: handle error
		}
	}
}

func sendSlackMessage(channelID, message string) error {
	msg := slack.MsgOptionText(message, false)

	_, _, err := slackBot.PostMessage(channelID, msg)
	if err != nil {
		return err
	}

	return nil
}
