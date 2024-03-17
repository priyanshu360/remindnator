package slackchannel

import (
	"testing"
)

var testChannelID = "C06PR53RUDC"

func TestSlackChannelSourceFetch(t *testing.T) {
	Init()
	sc, err := New(testChannelID)
	if err != nil {
		t.Errorf("Failed to get slack channel with id: %s, %v", testChannelID, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from Slack Channel: %v", err)
	}

}
