package user

import "context"


type GatewayImpl struct {
	client *Client
}

func NewGateway(client *Client) Gateway {
	return &GatewayImpl{
		client: client,
	}
}

func (g *GatewayImpl) GetUser(ctx context.Context, id int64) (string, error) {
	return g.client.GetUser(ctx, id)
}