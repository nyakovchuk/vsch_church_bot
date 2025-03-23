package service

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"

type DistanceManager interface {
	Distance(p1, p2 model.Coordinates) float64
}
