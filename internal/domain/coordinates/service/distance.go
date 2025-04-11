package service

import (
	"sort"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

type DistanceService interface {
	GetChurchesNearby(coords model.Coordinates, radius int, churches []church.Church) []church.DtoTelegram
}

type distanceService struct {
	repo repository.DistanceRepository
}

func NewDistanceService(repo repository.DistanceRepository) DistanceService {
	return &distanceService{repo: repo}
}

// Шукає найближчі церкви в радіусі Х метрів
func (s *distanceService) GetChurchesNearby(coords model.Coordinates, radius int, churches []church.Church) []church.DtoTelegram {

	if err := coords.Validate(); err != nil {
		return nil
	}

	findChurches := make([]church.DtoTelegram, 0)
	for _, church := range churches {
		churchCoords := model.Coordinates{
			Latitude:  church.Coordinates.Latitude,
			Longitude: church.Coordinates.Longitude,
		}
		distance := s.distance(coords, churchCoords)
		if distance <= float64(radius) {
			churchDtoTg := church.ToTelegramDto()
			churchDtoTg.Distance = distance

			findChurches = append(findChurches, churchDtoTg)
		}
	}

	return sortByDistance(findChurches)
}

func (s *distanceService) distance(coordinates1, coordinates2 model.Coordinates) float64 {

	distance := s.repo.Distance(coordinates1, coordinates2)

	return distance
}

func sortByDistance(churches []church.DtoTelegram) []church.DtoTelegram {
	sort.Slice(churches, func(i, j int) bool {
		return churches[i].Distance < churches[j].Distance
	})

	return churches
}
