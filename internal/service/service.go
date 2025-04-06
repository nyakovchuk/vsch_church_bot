package service

import (
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
)

type Service struct {
	Coordinates coordinates.CoordinatesService
	Distance    coordinates.DistanceService
	TgUser      tgUser.Service
	User        user.Service
}

func New(repo *repository.Repository) *Service {
	distanceService := coordinates.NewDistanceService(repo.DistanceRepository)
	coordinatesService := coordinates.NewCoordinatesService(repo.CoordinatesRepository)
	tgUserService := tgUser.NewService(repo.TgUserRepository)
	userService := user.NewService(repo.UserRepository)
	return &Service{
		Distance:    distanceService,
		Coordinates: coordinatesService,
		TgUser:      tgUserService,
		User:        userService,
	}
}
