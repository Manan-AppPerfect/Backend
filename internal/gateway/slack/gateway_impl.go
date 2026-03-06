package slack

import "context"

type GatewayImpl struct {
	client *Client
}

func NewGateway(client *Client) Gateway {
	return &GatewayImpl{
		client: client,
	}
}

func (g *GatewayImpl) SendMessage(ctx context.Context, message string) error {
	return g.client.SendMessage(ctx, message)
}