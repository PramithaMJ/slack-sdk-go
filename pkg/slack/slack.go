package slack

import (
    "context"
    "time"
)

type Config struct {
    Token           string
    BaseURL         string
    Timeout         time.Duration
    RetryAttempts   int
    RetryWaitTime   time.Duration
}

type SlackSDK struct {
    config Config
}

// NewSlackSDK initializes the Slack SDK with default configuration.
func NewSlackSDK(token string) *SlackSDK {
    return NewSlackSDKWithConfig(Config{
        Token:           token,
        BaseURL:         "https://slack.com/api/",
        Timeout:         30 * time.Second,
        RetryAttempts:   3,
        RetryWaitTime:   2 * time.Second,
    })
}

// NewSlackSDKWithConfig initializes the Slack SDK with custom configuration.
func NewSlackSDKWithConfig(config Config) *SlackSDK {
    return &SlackSDK{
        config: config,
    }
}
}
