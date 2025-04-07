package dto

import (
	"time"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
)

type RepositoryCoordinates struct {
	ID        int       `db:"id"`
	TgUserID  int64     `db:"tg_user_id"`
	Latitude  float64   `db:"latitude"`
	Longitude float64   `db:"longitude"`
	IsOnText  bool      `db:"is_on_text"`
	CreatedAt time.Time `db:"created_at"`
}

func ToModel(repoCoords RepositoryCoordinates) model.Coordinates {
	return model.Coordinates{
		ID:        repoCoords.ID,
		TgUserId:  repoCoords.TgUserID,
		Latitude:  repoCoords.Latitude,
		Longitude: repoCoords.Longitude,
		IsOnText:  repoCoords.IsOnText,
	}
}

func CoordinatesToModel(lat, lon float64) model.Coordinates {
	return model.Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}
}

func ModelToCoordinates(coords model.Coordinates) (lat, lon float64) {
	return coords.Latitude, coords.Longitude
}
