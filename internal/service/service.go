package service

import (
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
)

type Service struct {
	CoordinatesService coordinates.CoordinatesService
	DistanceService    coordinates.DistanceService
}

func New(repo *repository.Repository) *Service {
	distanceService := coordinates.NewDistanceService(repo.CoordinatesRepository)
	coordinatesService := coordinates.NewCoordinatesService()
	return &Service{
		DistanceService:    *distanceService,
		CoordinatesService: *coordinatesService,
	}
}
