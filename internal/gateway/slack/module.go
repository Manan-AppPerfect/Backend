package slack

import "go.uber.org/fx"

var SlackModule = fx.Options(
	fx.Provide(
		func() string {
			return "dummy-token" // Fx will use this string for Slack client
		},
		NewClient,
		NewGateway,
	),
)