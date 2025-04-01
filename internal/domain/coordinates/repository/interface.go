package repository

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"

type DistanceRepository interface {
	Distance(p1, p2 model.Coordinates) float64
}
