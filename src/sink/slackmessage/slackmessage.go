package slackmessage

import (
	"log"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/priyanshu360/remindnator/src/config"
	"github.com/priyanshu360/remindnator/src/event"

	"github.com/slack-go/slack"
)

type sink struct {
	channelId     string
	crontab       string
	lastPublished time.Time
}

var slackBot *slack.Client

func Init() error {
	slackBot = slack.New(config.SLACK_TOKEN)
	return nil
}

func New(id string, crontab string) *sink {
	return &sink{
		channelId:     id,
		crontab:       crontab,
		lastPublished: time.Now().Add(-1 * 24 * time.Hour),
	}
}

func (s *sink) Publish(event []event.Event) {
	currentTime := time.Now()
	nextPubTime := cronexpr.MustParse(s.crontab).Next(s.lastPublished)
	if currentTime.After(nextPubTime) {
		for _, e := range event {
			if err := sendSlackMessage(s.channelId, e.String()); err != nil {
				log.Printf("Error sendSlackMessage for channel %v : %v", s.channelId, err)
				continue
			}
		}
		s.lastPublished = currentTime
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
