package service

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

type DistanceService interface {
	GetChurchesNearby(latitude, longitude, radius float64) (int, error)
}

type distanceService struct {
	repo repository.DistanceRepository
}

// NewShopService создает новый сервис магазинов.
func NewDistanceService(repo repository.DistanceRepository) DistanceService {
	return &distanceService{repo: repo}
}

func (s *distanceService) GetChurchesNearby(latitude, longitude, radius float64) (int, error) {

	coordinates := dto.CoordinatesToModel(latitude, longitude)
	if err := coordinates.Validate(); err != nil {
		return 0, err
	}

	// получить координаты церквей (отдельный сервис)

	// churches := make([]string, 10)
	// // получить расстояние между координатами
	// for church := range churches {
	// 	if s.distance(coordinates, church.coordinates) < radius {

	// 		// добавить церкву в список
	// 	}
	// }

	return 0, nil

}

// GetShopsNearby ищет магазины в радиусе X км.
func (s *distanceService) distance(coordinates1, coordinates2 model.Coordinates) (float64, error) {

	distance := s.repo.Distance(coordinates1, coordinates2)

	return distance, nil
}
