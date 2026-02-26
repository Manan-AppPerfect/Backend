package team

import (
	"github.com/Manan-AppPerfect/Backend/internal/team/repository"
	"github.com/Manan-AppPerfect/Backend/internal/team/service"
	"go.uber.org/fx"
)

var TeamModule = fx.Options(
	fx.Provide(
		repository.NewRepo,
		service.NewService,
	),
)