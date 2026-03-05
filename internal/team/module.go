package team

import (
	"github.com/Manan-AppPerfect/Backend/internal/common/repository"
	"github.com/Manan-AppPerfect/Backend/internal/models"
	"github.com/Manan-AppPerfect/Backend/internal/team/service"
	"go.uber.org/fx"
)

var TeamModule = fx.Options(
	fx.Provide(
		repository.NewMemoryRepository[models.Team],
		service.NewService,
	),
)