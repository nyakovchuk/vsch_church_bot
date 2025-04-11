package repository

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/paulmach/orb/geo"
)

type DistanceRepository interface {
	Distance(p1, p2 model.Coordinates) float64
}

type coordinatesOrbRepository struct{}

func NewOrbRepository() DistanceRepository {
	return &coordinatesOrbRepository{}
}

func (c *coordinatesOrbRepository) Distance(p1, p2 model.Coordinates) float64 {

	p1DTO := dto.ModelToOrb(p1)
	p2DTO := dto.ModelToOrb(p2)

	// HaversineDistance
	distance := geo.Distance(p1DTO, p2DTO)

	return distance
}
