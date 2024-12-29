package slack

type MessagePayload struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

func (s *SlackSDK) SendMessage(channel, text string) (map[string]interface{}, error) {
	payload := MessagePayload{
		Channel: channel,
		Text:    text,
	}
	return s.sendRequest("chat.postMessage", "POST", payload)
}