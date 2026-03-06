package user

import "go.uber.org/fx"

var GatewayModule = fx.Options(
	fx.Provide(
		func() *Client {
			return NewClient("http://localhost:9000")
		},
		NewGateway,
	),
)