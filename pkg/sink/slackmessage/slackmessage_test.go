package slackmessage

import (
	"log"
	"testing"
	"time"

	"github.com/priyanshu360/remindnator/internal/event"
	"github.com/priyanshu360/remindnator/util"
)

var (
	testChannelID = "C06PR53RUDC"
	testCron      = "* * * * *"
)

func TestSlackMessageSinkPublish(t *testing.T) {
	if err := util.LoadLocalEnvFile(); err != nil {
		log.Println("Can't run tests local.env not found")
		return
	}
	Init()

	sc := New(testChannelID, testCron)
	testEventsList := []event.Event{event.New("first", time.Now(), false), event.New("second", time.Now(), true)}
	sc.Publish(testEventsList)
}
