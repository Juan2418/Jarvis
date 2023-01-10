package messages

import (
	"bytes"
)

type SlackWebhookMessage struct {
	Text string `json:"text"`
}

type SlackDirectMessage struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

type ListMessage struct {
	Items []string
}

func (m *ListMessage) String() string {
	var buffer bytes.Buffer
	for _, item := range m.Items {
		buffer.WriteString("\n- " + item)
	}
	return buffer.String()
}
