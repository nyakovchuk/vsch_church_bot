package dto

import (
	"time"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
)

type RepositoryCoordinates struct {
	ID         int       `db:"id"`
	PlatformID int       `db:"platform_id"`
	ExternalID string    `db:"external_id"`
	Latitude   float64   `db:"latitude"`
	Longitude  float64   `db:"longitude"`
	IsOnText   bool      `db:"is_on_text"`
	CreatedAt  time.Time `db:"created_at"`
}

func (repoCoords RepositoryCoordinates) ToModel() model.Coordinates {
	return model.Coordinates{
		ID:         repoCoords.ID,
		PlatformID: repoCoords.PlatformID,
		ExternalID: repoCoords.ExternalID,
		Latitude:   repoCoords.Latitude,
		Longitude:  repoCoords.Longitude,
		IsOnText:   repoCoords.IsOnText,
	}
}
