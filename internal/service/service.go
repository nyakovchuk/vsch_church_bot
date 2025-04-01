package service

import (
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
)

type Service struct {
	Coordinates coordinates.CoordinatesService
	Distance    coordinates.DistanceService
}

func New(repo *repository.Repository) *Service {
	distanceService := coordinates.NewDistanceService(repo.DistanceRepository)
	coordinatesService := coordinates.NewCoordinatesService()
	return &Service{
		Distance:    *distanceService,
		Coordinates: *coordinatesService,
	}
}
