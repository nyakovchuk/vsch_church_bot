package service

import (
	"sort"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

type DistanceService interface {
	GetChurchesNearby(coords model.Coordinates, radius int, churches []church.Church) []church.DtoResponse
	FindTopNNearestChurches(coords model.Coordinates, topN int, churches []church.Church) []church.DtoResponse
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

func (s *distanceService) FindTopNNearestChurches(coords model.Coordinates, topN int, churches []church.Church) []church.DtoResponse {

	churchesWithDistances := make([]*church.ChurchWithDistance, 0, len(churches))
	// Вычисляем расстояние для каждой церкви
	for i := range churches {
		churchCoords := model.GeoToModel(
			churches[i].Coordinates.Latitude,
			churches[i].Coordinates.Longitude,
		)
		cwd := &church.ChurchWithDistance{
			Church:   &churches[i],
			Distance: s.distance(coords, churchCoords),
		}
		churchesWithDistances = append(churchesWithDistances, cwd)
	}

	// Создаём слайс для топ-N
	topNChurches := make([]*church.ChurchWithDistance, 0, topN)

	// Поиск топ-N ближайших
	for i := range churchesWithDistances {
		c := churchesWithDistances[i]

		// Вставка на нужную позицию
		inserted := false
		for j := 0; j < len(topNChurches); j++ {
			if c.Distance < topNChurches[j].Distance {
				topNChurches = append(topNChurches[:j+1], topNChurches[j:]...)
				topNChurches[j] = c
				inserted = true
				break
			}
		}

		// Если не вставили и ещё есть место — добавляем в конец
		if !inserted && len(topNChurches) < topN {
			topNChurches = append(topNChurches, c)
		}

		// Ограничиваем длину до N
		if len(topNChurches) > topN {
			topNChurches = topNChurches[:topN]
		}
	}

	findChurches := make([]church.DtoResponse, 0, topN)
	for _, church := range topNChurches {
		findChurches = append(findChurches, church.ToDtoResponse())
	}

	return findChurches
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
