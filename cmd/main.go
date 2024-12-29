package main

import (
	"fmt"
	"log"

	"github.com/pramithamj/slack-sdk-go/pkg/slack"
)

func main() {
	token := "xoxb-your-slack-bot-token"
	sdk := slack.NewSlackSDK(token)

	channel := "C123456789"
	message := "Hello, Slack from Go SDK!"

	response, err := sdk.SendMessage(channel, message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	fmt.Printf("Message sent successfully: %v\n", response)
}