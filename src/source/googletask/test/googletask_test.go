package test

import (
	"testing"

	"github.com/priyanshu360/remindnator/src/source/googletask"
)

var testTaskTitle = "My Tasks"

func TestGoogleTaskSourceFetch(t *testing.T) {
	googletask.Init()

	sc, err := googletask.New(testTaskTitle)
	if err != nil {
		t.Errorf("Failed to get google task with title: %s, %v", testTaskTitle, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from google task: %v", err)
	}

}
