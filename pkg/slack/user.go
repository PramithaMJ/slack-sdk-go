package slack

import "context"

func (s *SlackSDK) GetUserInfo(ctx context.Context, userID string) (map[string]interface{}, error) {
	endpoint := "users.info"
	payload := map[string]string{"user": userID}
	return s.sendRequest(ctx, endpoint, "GET", payload)
}