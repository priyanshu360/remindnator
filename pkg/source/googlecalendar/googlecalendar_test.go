package googlecalendar

import (
	"testing"
)

var testCalanderID = "priyanshurajput360@gmail.com"

func TestGoogleCalendarSourceFetch(t *testing.T) {
	Init()

	sc, err := New(testCalanderID)
	if err != nil {
		t.Errorf("Failed to get google cal with id: %s, %v", testCalanderID, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from google cal: %v", err)
	}
}
