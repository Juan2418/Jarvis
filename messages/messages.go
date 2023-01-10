package messages

import (
	"encoding/json"
	"jarvis/slack"
	"jarvis/utils"
	"sync"

	_ "github.com/joho/godotenv/autoload"
)

func SendSlackWebhookMessage(messageBody string) {
	message := SlackWebhookMessage{
		Text: messageBody,
	}
	marshalledMessage, _ := json.Marshal(message)

	slack.SendWebhookMessage(marshalledMessage)
}

func SendSlackDM(userId string, messageBody string) {
	message := SlackDirectMessage{
		Channel: userId,
		Text:    messageBody,
	}

	marshalledMessage, _ := json.Marshal(message)
	slack.SendDM(marshalledMessage)
}

func SendSlackDMBatch(userIds []string, messageBody string) {
	syncBuf := utils.GetSyncBuffer()
	var wg sync.WaitGroup
	wg.Add(len(userIds))

	for _, userId := range userIds {
		<-syncBuf
		go func(userId string) {
			SendSlackDM(userId, messageBody)
			defer wg.Done()
			syncBuf <- true
		}(userId)
	}

	wg.Wait()
}
