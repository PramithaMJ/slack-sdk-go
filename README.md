Here’s the entire content as a single .md file:

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

Enjoy coding with the Slack SDK for Go! 🧑‍💻
