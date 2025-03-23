package model

import "errors"

var (
	ErrCorrectLatitude  = errors.New("некорректная широта (должна быть в диапазоне -90...90)")
	ErrCorrectLongitude = errors.New("некорректная долгота (должна быть в диапазоне -180...180)")
)

type Coordinates struct {
	Latitude  float64
	Longitude float64
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
