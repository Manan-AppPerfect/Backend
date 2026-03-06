package slack

import "go.uber.org/fx"

var SlackModule = fx.Options(
	fx.Provide(
		NewClient,
		NewGateway,
	),
)