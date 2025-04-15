package service

import (
	"sort"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

type DistanceService interface {
	GetChurchesNearby(coords model.Coordinates, radius int, churches []church.Church) []church.DtoResponse
}

type distanceService struct {
	repo repository.DistanceRepository
}

func NewDistanceService(repo repository.DistanceRepository) DistanceService {
	return &distanceService{repo: repo}
}

// Шукає найближчі церкви в радіусі Х метрів
func (s *distanceService) GetChurchesNearby(coords model.Coordinates, radius int, churches []church.Church) []church.DtoResponse {

	if err := coords.Validate(); err != nil {
		return nil
	}

	findChurches := make([]church.DtoResponse, 0)
	for _, church := range churches {
		churchCoords := model.GeoToModel(
			church.Coordinates.Latitude,
			church.Coordinates.Longitude,
		)

		distance := s.distance(coords, churchCoords)
		if distance <= float64(radius) {
			churchDtoResp := church.ToDtoResponse()
			churchDtoResp.Distance = distance

			findChurches = append(findChurches, churchDtoResp)
		}
	}

	return sortByDistance(findChurches)
}

func (s *distanceService) distance(coordinates1, coordinates2 model.Coordinates) float64 {

	distance := s.repo.Distance(coordinates1, coordinates2)

	return distance
}

func sortByDistance(churches []church.DtoResponse) []church.DtoResponse {
	sort.Slice(churches, func(i, j int) bool {
		return churches[i].Distance < churches[j].Distance
	})

	return churches
}
