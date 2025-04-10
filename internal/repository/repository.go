package repository

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/repository"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/user"
)

type Repository struct {
	DistanceRepository    repository.DistanceRepository
	CoordinatesRepository repository.CoordinatesRepository
	TgUserRepository      tgUser.Repository
	UserRepository        user.Repository
	ChurchRepository      church.Repository
}

func New(db *sql.DB) *Repository {
	return &Repository{
		DistanceRepository:    repository.NewOrbRepository(),
		CoordinatesRepository: repository.NewCoordinatesRepository(db),
		TgUserRepository:      tgUser.NewRepository(db),
		UserRepository:        user.NewRepository(db),
		ChurchRepository:      church.NewRepository(db),
	}
}
