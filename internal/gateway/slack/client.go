package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type Client struct {
	webhookURL string
	httpClient *http.Client
}

func NewClient(webhookURL string) *Client {
	return &Client{
		webhookURL: webhookURL,
		httpClient: &http.Client{},
	}
}

type slackPayload struct {
	Text string `json:"text"`
}

func (c *Client) SendMessage(ctx context.Context, message string) error {

	payload := slackPayload{
		Text: message,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return err 
	}

	req, err := http.NewRequestWithContext(
		ctx, 
		http.MethodPost,
		c.webhookURL,
		bytes.NewBuffer(body),
	)

	if err != nil {
		return err 
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err 
	}

	defer resp.Body.Close()
	
	return nil
}