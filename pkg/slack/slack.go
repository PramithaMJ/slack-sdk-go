package slack

type SlackSDK struct {
	Token   string
	BaseURL string
}

// NewSlackSDK initializes the Slack SDK.
func NewSlackSDK(token string) *SlackSDK {
	return &SlackSDK{
		Token:   token,
		BaseURL: "https://slack.com/api/",
	}
}