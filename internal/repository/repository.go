package repository

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
)

type Repository struct {
	DistanceRepository    repository.DistanceRepository
	CoordinatesRepository repository.CoordinatesRepository
	UserRepository        user.Repository
	ChurchRepository      church.Repository
	PlatformRepository    platform.Repository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DistanceRepository:    repository.NewOrbRepository(),
		CoordinatesRepository: repository.NewCoordinatesRepository(db),
		UserRepository:        user.NewRepository(db),
		ChurchRepository:      church.NewRepository(db),
		PlatformRepository:    platform.NewRepository(db),
	}
}
