package googletask

import (
	"testing"
)

var testTaskTitle = "My Tasks"

func TestGoogleTaskSourceFetch(t *testing.T) {
	Init()

	sc, err := New(testTaskTitle)
	if err != nil {
		t.Errorf("Failed to get google task with title: %s, %v", testTaskTitle, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from google task: %v", err)
	}

}
