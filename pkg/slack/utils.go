package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// SlackError represents an error response from the Slack API
type SlackError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Ok      bool   `json:"ok"`
}

// Error implements the error interface
func (e *SlackError) Error() string {
	return fmt.Sprintf("Slack API error: %s - %s", e.Code, e.Message)
}

func (s *SlackSDK) sendRequest(ctx context.Context, endpoint, method string, payload interface{}) (map[string]interface{}, error) {
    if ctx == nil {
        ctx = context.Background()
    }

    url := s.config.BaseURL + endpoint
    body, err := json.Marshal(payload)
    if err != nil {
        return nil, fmt.Errorf("failed to marshal payload: %w", err)
    }

    req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
    if err != nil {
        return nil, fmt.Errorf("failed to create request: %w", err)
    }

    req.Header.Set("Authorization", "Bearer "+s.config.Token)
    req.Header.Set("Content-Type", "application/json; charset=utf-8")

    client := &http.Client{
        Timeout: s.config.Timeout,
    }

    var lastErr error
    for attempt := 0; attempt <= s.config.RetryAttempts; attempt++ {
        if attempt > 0 {
            select {
            case <-ctx.Done():
                return nil, ctx.Err()
            case <-time.After(s.config.RetryWaitTime):
            }
        }

        resp, err := client.Do(req)
        if err != nil {
            lastErr = fmt.Errorf("request failed: %w", err)
            continue
        }

        defer resp.Body.Close()
        var response map[string]interface{}

        if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
            lastErr = fmt.Errorf("failed to decode response: %w", err)
            continue
        }

        // Check Slack's ok field
        if ok, _ := response["ok"].(bool); !ok {
            errorMsg := "unknown error"
            if msg, ok := response["error"].(string); ok {
                errorMsg = msg
            }
            return nil, &SlackError{
                Code:    resp.Status,
                Message: errorMsg,
                Ok:      false,
            }
        }

        return response, nil
    }

    return nil, fmt.Errorf("max retries exceeded: %w", lastErr)
}
