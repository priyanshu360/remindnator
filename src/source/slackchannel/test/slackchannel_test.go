package test

import (
	"testing"

	"github.com/priyanshu360/remindnator/src/source/slackchannel"
)

var testChannelID = "C06PR53RUDC"

func TestSlackChannelSourceFetch(t *testing.T) {
	slackchannel.Init()

	sc, err := slackchannel.New(testChannelID)
	if err != nil {
		t.Errorf("Failed to get slack channel with id: %s, %v", testChannelID, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from Slack Channel: %v", err)
	}

}
