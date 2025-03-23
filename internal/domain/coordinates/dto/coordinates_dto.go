package dto

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
)

func CoordinatesToModel(lat, lon float64) model.Coordinates {
	return model.Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}
}

func ModelToCoordinates(coords model.Coordinates) (lat, lon float64) {
	return coords.Latitude, coords.Longitude
}
