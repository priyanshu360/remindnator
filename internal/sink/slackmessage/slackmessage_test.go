package slackmessage

var (
	testChannelID = "C06PR53RUDC"
	testCron      = "* * * * *"
)

// func TestSlackMessageSinkPublish(t *testing.T) {
// 	if err := util.LoadLocalEnvFile(); err != nil {
// 		log.Println("Can't run tests local.env not found")
// 		return
// 	}
// 	Init()

// 	sc := New(testChannelID, testCron)
// 	start := time.Now()
// 	end := time.Now().Add(time.Hour)
// 	testEventsList := []event.Event{event.New("first", &start, &end, false)}
// 	sc.Publish(testEventsList)
// }
