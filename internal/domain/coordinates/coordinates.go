package coordinates

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
)

func NewDistanceService() service.DistanceService {
	repository := repository.NewOrbRepository()

	return service.NewDistanceService(repository)
}

// func NewService() service.CoordinatesService {
// 	repository := repository.NewCoordinatesRepository()
// 	return service.NewCoordinatesService()
// }
