package model

import "errors"

var (
	ErrCorrectLatitude  = errors.New("некорректная широта (должна быть в диапазоне -90...90)")
	ErrCorrectLongitude = errors.New("некорректная долгота (должна быть в диапазоне -180...180)")
)

type Coordinates struct {
	ID         int
	PlatformID int
	ExternalID string
	Latitude   float64
	Longitude  float64
	IsOnText   bool
}

func ToCoordinates(platformId int, externalId string, latitude, longitude float64, isOnText bool) Coordinates {
	return Coordinates{
		PlatformID: platformId,
		ExternalID: externalId,
		Latitude:   latitude,
		Longitude:  longitude,
		IsOnText:   isOnText,
	}
}

func GeoToModel(lat, lon float64) Coordinates {
	return Coordinates{
		Latitude:  lat,
		Longitude: lon,
	}
}

func ModelToGeo(coords Coordinates) (lat, lon float64) {
	return coords.Latitude, coords.Longitude
}

func (c *Coordinates) Validate() error {
	if c.Latitude < -90 || c.Latitude > 90 {
		return ErrCorrectLatitude
	}
	if c.Longitude < -180 || c.Longitude > 180 {
		return ErrCorrectLongitude
	}
	return nil
}
