package slack

type MessagePayload struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
}

import "context"

// SendMessage sends a simple text message to a channel
func (s *SlackSDK) SendMessage(ctx context.Context, channel, text string) (map[string]interface{}, error) {
    if channel == "" {
        return nil, errors.New("channel cannot be empty")
    }
    
    payload := MessagePayload{
        Channel: channel,
        Text:    text,
    }
    return s.sendRequest(ctx, "chat.postMessage", "POST", payload)
}

// SendRichMessage sends a message with blocks and attachments
func (s *SlackSDK) SendRichMessage(ctx context.Context, payload MessagePayload) (map[string]interface{}, error) {
    if payload.Channel == "" {
        return nil, errors.New("channel cannot be empty")
    }
    
    if payload.Text == "" && len(payload.Blocks) == 0 && len(payload.Attachments) == 0 {
        return nil, errors.New("message must contain either text, blocks, or attachments")
    }
    
    return s.sendRequest(ctx, "chat.postMessage", "POST", payload)
}

// SendThreadReply sends a message as a reply in a thread
func (s *SlackSDK) SendThreadReply(ctx context.Context, channel, threadTs, text string) (map[string]interface{}, error) {
    if channel == "" || threadTs == "" {
        return nil, errors.New("channel and thread_ts cannot be empty")
    }
    
    payload := MessagePayload{
        Channel:  channel,
        Text:     text,
        ThreadTS: threadTs,
    }
    return s.sendRequest(ctx, "chat.postMessage", "POST", payload)
}
