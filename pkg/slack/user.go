package slack

func (s *SlackSDK) GetUserInfo(userID string) (map[string]interface{}, error) {
	endpoint := "users.info"
	payload := map[string]string{"user": userID}
	return s.sendRequest(endpoint, "GET", payload)
}