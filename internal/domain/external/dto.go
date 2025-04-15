package external

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"

type ExternalRepository struct {
	ID           string
	PlatformID   int
	PlatformName string
}

func (e ExternalRepository) ToModel() External {
	return External{
		ID:       e.ID,
		Platform: platform.Platform{ID: e.PlatformID, Name: e.PlatformName},
	}
}
