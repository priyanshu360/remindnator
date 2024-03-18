package slackchannel

import (
	"log"
	"testing"

	"github.com/priyanshu360/remindnator/util"
)

var testChannelID = "C06PR53RUDC"

func TestSlackChannelSourceFetch(t *testing.T) {
	if err := util.LoadLocalEnvFile(); err != nil {
		log.Println("Can't run tests local.env not found")
		return
	}
	Init()
	sc, err := New(testChannelID)
	if err != nil {
		t.Errorf("Failed to get slack channel with id: %s, %v", testChannelID, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from Slack Channel: %v", err)
	}

}
