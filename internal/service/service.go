package service

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	coordinates "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/service"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/language"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
	"github.com/nyakovchuk/vsch_church_bot/internal/repository"
)

type Service struct {
	Coordinates coordinates.CoordinatesService
	Distance    coordinates.DistanceService
	User        user.Service
	Church      church.Service
	Platform    platform.Service
	Language    language.Service
}

func New(repo *repository.Repository) *Service {
	distanceService := coordinates.NewDistanceService(repo.DistanceRepository)
	coordinatesService := coordinates.NewCoordinatesService(repo.CoordinatesRepository)
	userService := user.NewService(repo.UserRepository)
	churchService := church.NewService(repo.ChurchRepository)
	platformService := platform.NewService(repo.PlatformRepository)
	languageService := language.NewService(repo.LanguageRepository)

	return &Service{
		Distance:    distanceService,
		Coordinates: coordinatesService,
		User:        userService,
		Church:      churchService,
		Platform:    platformService,
		Language:    languageService,
	}
}
