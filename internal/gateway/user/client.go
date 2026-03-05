package user

import (
	"context"
)

type Client struct {
	baseURL string
}

func NewClient() *Client {
	return &Client{
		baseURL: "http://localhost:9000",
	}
}

func (c *Client) GetUser(ctx context.Context, id int64) (string, error) {
	println("Calling user service......")
	return "Manan", nil
}