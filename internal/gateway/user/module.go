package user

import "go.uber.org/fx"

var GatewayModule = fx.Options(
	fx.Provide(
		NewClient,
		NewGateway,
	),
)