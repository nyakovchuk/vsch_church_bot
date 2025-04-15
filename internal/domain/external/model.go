package external

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"

type External struct {
	ID       string
	Platform platform.Platform
}

func ToModel(id string, platform platform.Platform) External {
	return External{
		ID:       id,
		Platform: platform,
	}
}

func (e External) ToRepository() ExternalRepository {
	return ExternalRepository{
		ID:           e.ID,
		PlatformID:   e.Platform.ID,
		PlatformName: e.Platform.Name,
	}
}
