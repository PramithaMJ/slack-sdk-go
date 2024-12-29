package tests

import (
	"testing"

	"github.com/pramithamj/slack-sdk-go/pkg/slack"
)

func TestSendMessage(t *testing.T) {
	mockSDK := slack.NewSlackSDK("mock-token")
	_, err := mockSDK.SendMessage("C123456789", "Test message")
	if err != nil {
		t.Errorf("SendMessage failed: %v", err)
	}
}