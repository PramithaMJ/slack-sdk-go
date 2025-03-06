package tests

import (
	"context"
	"testing"

	"github.com/pramithamj/slack-sdk-go/pkg/slack"
)

func TestSendMessage(t *testing.T) {
	mockSDK := slack.NewSlackSDK("mock-token")
	_, err := mockSDK.SendMessage(context.Background(), "C123456789", "Test message")
	if err != nil {
		t.Errorf("SendMessage failed: %v", err)
	}
}