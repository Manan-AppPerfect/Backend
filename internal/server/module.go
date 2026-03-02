package server

import (
	"go.uber.org/fx"
)

//fx.Options lets you group things together


var ServerModule = fx.Options(
	fx.Provide(
		NewTeamHandler,
	),
	fx.Invoke(
		NewHTTPServer,
	),
)