package external

import "github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"

type External struct {
	Id       string
	Platform platform.Platform
}

func ToModel(id string, platform platform.Platform) External {
	return External{
		Id:       id,
		Platform: platform,
	}
}
