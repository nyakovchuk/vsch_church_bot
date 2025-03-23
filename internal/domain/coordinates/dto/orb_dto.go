package dto

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/paulmach/orb"
)

func OrbToModel(p orb.Point) model.Coordinates {
	return model.Coordinates{
		Latitude:  p.Point().Lat(),
		Longitude: p.Point().Lon(),
	}
}

func ModelToOrb(coords model.Coordinates) orb.Point {
	return orb.Point{
		coords.Longitude,
		coords.Latitude,
	}
}
