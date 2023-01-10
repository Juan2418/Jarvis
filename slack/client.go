package slack

import (
	"bytes"
	"encoding/json"
	"jarvis/debugging"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var WEBHOOK_URL = os.Getenv("SLACK_WEBHOOK_URL")
var SLACK_API_TOKEN = os.Getenv("SLACK_API_TOKEN")
var SLACK_API_BASE_URL = os.Getenv("SLACK_API_BASE_URL")

func SendRequest(url string, payload []byte) error {
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
	}

	client := &http.Client{}
	result, err := client.Do(request)

	if err != nil {
		log.Println(err)
	}

	log.Println(result.Status)

	return err
}

func SendWebhookMessage(payload []byte) error {
	return SendRequest(WEBHOOK_URL, payload)
}

func DoSlackApiRequest(method string, endpoint string, payload []byte) (*http.Response, error) {
	url := SLACK_API_BASE_URL + endpoint
	request, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		log.Println(err)
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+SLACK_API_TOKEN)

	client := &http.Client{}
	result, err := client.Do(request)

	debugging.LogCurl(request)

	if err != nil {
		log.Println(err)
	}

	log.Println(result.Status)

	return result, err
}

func Get(endpoint string) (*http.Response, error) {
	return DoSlackApiRequest(http.MethodGet, endpoint, nil)
}

func Post(endpoint string, payload []byte) (*http.Response, error) {
	return DoSlackApiRequest(http.MethodPost, endpoint, payload)
}

func GetSlackIDByEmail(email string) SlackIDByEmailResponse {
	response, err := Get("users.lookupByEmail?email=" + email)

	if err != nil {
		log.Println(err)
	}

	var slackIDByEmailResponse SlackIDByEmailResponse
	err = json.NewDecoder(response.Body).Decode(&slackIDByEmailResponse)

	if err != nil {
		log.Println(err)
	}

	return slackIDByEmailResponse
}

func SendDM(payload []byte) error {
	_, err := Post("chat.postMessage", payload)

	return err
}
