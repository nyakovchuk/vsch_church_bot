package coordinates

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
)

func New() *service.DistanceService {
	repository := repository.NewOrbRepository()

	return service.NewDistanceService(repository)
}
