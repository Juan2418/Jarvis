package main

import (
	"jarvis/messages"
	"jarvis/slack"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	isDebugging := os.Getenv("DEBUG") == "true"

	if isDebugging {
		testList := messages.ListMessage{
			Items: []string{
				"Item 1",
				"Item 2",
				"Item 3",
			},
		}
		testMessage := "This is a test message" + testList.String()
		testEmail := os.Getenv("TEST_EMAIL")

		slackResponse := slack.GetSlackIDByEmail(testEmail)

		var idList []string

		for i := 0; i < 10; i++ {
			idList = append(idList, slackResponse.User.ID)
		}

		messages.SendSlackDMBatch(idList, testMessage)
	}
}
