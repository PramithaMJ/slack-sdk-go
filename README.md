# 🚀 Slack SDK for Go  

A **lightweight** and **modular** SDK for building Slack bots and apps in **Go**. With this SDK, you can easily send messages, manage users, and extend functionality to integrate with the Slack API seamlessly.  

---

### 🌟 Features  

- 📨 **Send messages** to Slack channels  
- 👤 **Fetch user information**  
- ⚡️ **Extendable** for additional Slack APIs  
- 🧩 **Lightweight** and modular design  
- 🎯 **Easy integration** into Go projects  

---

### 🛠 Getting Started  

#### ✅ Prerequisites  

- **Go** 1.18 or later  
- A **Slack Bot Token** with the required permissions.  
  > [Create a Slack App](https://api.slack.com/apps) to generate your token.  

#### 📦 Installation  

Add the SDK to your project using:  

```sh  
go get github.com/pramithamj/slack-sdk-go/pkg/slack  
```
✨ Example Usage

Here’s a quick example to send a message using the SDK:
```go
package main  

import (  
	"fmt"  
	"log"  

	"github.com/pramithamj/slack-sdk-go/pkg/slack"  
)  

func main() {  
	// Replace with your actual Slack Bot Token  
	token := "xoxb-your-slack-bot-token"  
	sdk := slack.NewSlackSDK(token)  

	channel := "C123456789" // Slack channel ID  
	message := "Hello, Slack from Go SDK!"  

	response, err := sdk.SendMessage(channel, message)  
	if err != nil {  
		log.Fatalf("Error sending message: %v", err)  
	}  

	fmt.Printf("Message sent successfully: %v\n", response)  
}  
```
#### 📖 API Methods

1️⃣ SendMessage

Send a message to a Slack channel.

📥 Parameters:
	•	channel (string): The channel ID where the message will be sent.
	•	text (string): The message text.

📤 Example:
```go
response, err := sdk.SendMessage("C123456789", "Hello, Slack!")  
if err != nil {  
	log.Fatalf("Error: %v", err)  
}  
fmt.Println(response)  
```
2️⃣ GetUserInfo

Fetch details about a Slack user.

📥 Parameters:
	•	userID (string): The ID of the user.

📤 Example:
```go
response, err := sdk.GetUserInfo("U123456789")  
if err != nil {  
	log.Fatalf("Error: %v", err)  
}  
fmt.Println(response)  
```
#### 🔍 Testing

Run all tests using:

go test ./tests/...  

#### 🤝 Contributing

We welcome contributions! To contribute:
	1.	Fork the repository
	2.	Create a new branch: git checkout -b feature-name
	3.	Commit your changes: git commit -m 'Add feature'
	4.	Push the branch: git push origin feature-name
	5.	Open a Pull Request

#### 📜 License

This project is licensed under the MIT License. See the LICENSE file for more details.

👤 Author

PramithaMJ
	•	💻 [PramithaMJ]](https://github.com/pramithamj)
	•	📧 lpramithamj@gmail.com

# Slack SDK for Go 🚀

A powerful and flexible Go SDK for interacting with the Slack API, providing a simple yet comprehensive interface for sending messages and managing Slack communications.

## Features ✨

- 📝 Simple and rich message support
- 🧵 Thread reply functionality
- 🔄 Automatic retries with configurable attempts
- ⏱️ Context support for timeouts and cancellation
- 🎯 Custom error types and proper error handling
- 🔐 Secure configuration via environment variables
- 📦 Block kit and message attachments support

## Installation 📥

```bash
go get github.com/pramithamj/slack-go-sdk
```

## Quick Start 🚀

```go
package main

import (
    "context"
    "log"
    "os"
    "time"

    "github.com/pramithamj/slack-go-sdk/pkg/slack"
)

func main() {
    // Get token from environment
    token := os.Getenv("SLACK_BOT_TOKEN")
    if token == "" {
        log.Fatal("SLACK_BOT_TOKEN is required")
    }

    // Initialize SDK
    sdk := slack.NewSlackSDK(token)

    // Send a message
    ctx := context.Background()
    response, err := sdk.SendMessage(ctx, "CHANNEL_ID", "Hello from Go SDK!")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Message sent: %v", response)
}
```

## Advanced Usage 🔧

### Custom Configuration

```go
sdk := slack.NewSlackSDKWithConfig(slack.Config{
    Token:         token,
    BaseURL:       "https://slack.com/api/",
    Timeout:       10 * time.Second,
    RetryAttempts: 2,
    RetryWaitTime: time.Second,
})
```

### Rich Messages with Blocks

```go
richMessage := slack.MessagePayload{
    Channel: channelID,
    Blocks: []slack.Block{
        {
            Type: "section",
            Text: &slack.TextObject{
                Type: "mrkdwn",
                Text: "*Hello!*\nThis is a rich message.",
            },
        },
    },
}

response, err := sdk.SendRichMessage(ctx, richMessage)
```

### Thread Replies

```go
response, err := sdk.SendThreadReply(ctx, channelID, threadTS, "This is a reply!")
```

## Environment Variables 🔐

- `SLACK_BOT_TOKEN`: Your Slack bot token (required)
- `SLACK_CHANNEL_ID`: Default channel ID for messages

## Error Handling 🔍

```go
response, err := sdk.SendMessage(ctx, channelID, message)
if err != nil {
    if slackErr, ok := err.(*slack.SlackError); ok {
        log.Printf("Slack API error: %s", slackErr.Message)
    } else {
        log.Printf("Other error: %v", err)
    }
}
```

## Contributing 🤝

Contributions are welcome! Feel free to:
- Report bugs
- Suggest new features
- Submit PRs

## Contact 📫

- 💻 [PramithaMJ](https://github.com/pramithamj)
- 📧 lpramithamj@gmail.com

## License 📄

This project is licensed under the MIT License - see the LICENSE file for details.

---
Last Updated: March 1, 2025