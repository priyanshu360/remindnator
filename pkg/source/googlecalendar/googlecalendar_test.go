package googlecalendar

import (
	"fmt"
	"log"
	"testing"

	"github.com/priyanshu360/remindnator/util"
)

var testCalanderID = "priyanshurajput360@gmail.com"

func TestGoogleCalendarSourceFetch(t *testing.T) {
	if err := util.LoadLocalEnvFile(); err != nil {
		log.Println("Can't run tests local.env not found", err)
		return
	}

	if err := util.LoadTokenFromFile(); err != nil {
		fmt.Println("Can't run tests without a valid token.json file.")
		return
	}

	Init()

	sc, err := New(testCalanderID)
	if err != nil {
		t.Errorf("Failed to get google cal with id: %s, %v", testCalanderID, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from google cal: %v", err)
	}
}
