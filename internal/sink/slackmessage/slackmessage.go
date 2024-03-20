package slackmessage

import (
	"github.com/priyanshu360/remindnator/config"
	"github.com/priyanshu360/remindnator/internal/event"
	"github.com/priyanshu360/remindnator/util"

	"github.com/slack-go/slack"
)

type sink struct {
	channelId string
	id        string
}

var slackBot *slack.Client

func Init() error {
	slackBot = slack.New(config.SLACK_TOKEN)
	return nil
}

func New(id string) *sink {
	return &sink{
		channelId: id,
		id:        util.GenerateUUID(10),
	}
}

func (s *sink) Publish(e event.Event) error {
	return sendSlackMessage(s.channelId, e.String())
}

func sendSlackMessage(channelID, message string) error {
	msg := slack.MsgOptionText(message, false)

	_, _, err := slackBot.PostMessage(channelID, msg)
	if err != nil {
		return err
	}

	return nil
}
