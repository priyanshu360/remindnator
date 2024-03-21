package googletask

import (
	"fmt"
	"log"
	"testing"

	"github.com/priyanshu360/remindnator/util"
)

var testTaskTitle = "My Tasks"

func TestGoogleTaskSourceFetch(t *testing.T) {
	if err := util.LoadLocalEnvFile(); err != nil {
		log.Println("Can't run tests local.env not found")
		return
	}

	if err := util.LoadTokenFromFile(); err != nil {
		fmt.Println("Can't run TestGoogleTaskSourceFetch tests without a valid token.json file.")
		return
	}

	Init()

	sc, err := New(testTaskTitle)
	if err != nil {
		t.Errorf("Failed to get google task with title: %s, %v", testTaskTitle, err)
	}

	if err := sc.Fetch(); err != nil {
		t.Errorf("Failed to fetch from google task: %v", err)
	}

}
