package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "github.com/pramithamj/slack-sdk-go/pkg/slack"
)

func main() {
    // Get token from environment variable
    token := os.Getenv("SLACK_BOT_TOKEN")
    if token == "" {
        log.Fatal("SLACK_BOT_TOKEN environment variable is required")
    }

    // Initialize SDK with custom configuration
    sdk := slack.NewSlackSDKWithConfig(slack.Config{
        Token:         token,
        BaseURL:       "https://slack.com/api/",
        Timeout:       10 * time.Second,
        RetryAttempts: 2,
        RetryWaitTime: time.Second,
    })

    // Create context with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
    defer cancel()

    // Example 1: Send a simple message
    channel := os.Getenv("SLACK_CHANNEL_ID")
    if channel == "" {
        log.Fatal("SLACK_CHANNEL_ID environment variable is required")
    }

    response, err := sdk.SendMessage(ctx, channel, "Hello, Slack from Go SDK!")
    if err != nil {
        log.Printf("Error sending simple message: %v\n", err)
    } else {
        fmt.Printf("Simple message sent successfully: %v\n", response)
    }

    // Example 2: Send a rich message with blocks
    richMessage := slack.MessagePayload{
        Channel: channel,
        Blocks: []slack.Block{
            {
                Type: "section",
                Text: &slack.TextObject{
                    Type: "mrkdwn",
                    Text: "*Hello from Go SDK!*\nThis is a rich message with blocks.",
                },
            },
            {
                Type: "divider",
            },
        },
    }

    response, err = sdk.SendRichMessage(ctx, richMessage)
    if err != nil {
        log.Printf("Error sending rich message: %v\n", err)
    } else {
        fmt.Printf("Rich message sent successfully: %v\n", response)
    }
}
