package repository

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
)

type Repository struct {
	DistanceRepository    repository.DistanceRepository
	CoordinatesRepository repository.CoordinatesRepository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DistanceRepository:    repository.NewOrbRepository(),
		CoordinatesRepository: repository.NewCoordinatesRepository(db),
	}
}

// func (r *repo) Distance(p1, p2 model.Coordinates) float64 {
// 	return r.CoordinatesRepository.Distance(p1, p2)
// }
