package service

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
)

type Service struct {
	Coordinates coordinates.CoordinatesService
	Distance    coordinates.DistanceService
	TgUser      tgUser.Service
	User        user.Service
	Church      church.Service
	Platform    platform.Service
}

func New(repo *repository.Repository) *Service {
	distanceService := coordinates.NewDistanceService(repo.DistanceRepository)
	coordinatesService := coordinates.NewCoordinatesService(repo.CoordinatesRepository)
	tgUserService := tgUser.NewService(repo.TgUserRepository)
	userService := user.NewService(repo.UserRepository)
	churchService := church.NewService(repo.ChurchRepository)
	platformService := platform.NewService(repo.PlatformRepository)

	return &Service{
		Distance:    distanceService,
		Coordinates: coordinatesService,
		TgUser:      tgUserService,
		User:        userService,
		Church:      churchService,
		Platform:    platformService,
	}
}
