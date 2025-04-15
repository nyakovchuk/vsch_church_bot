package shareddata

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/church"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
)

type Data struct {
	Churches []church.Church
	Platform platform.Platform
}
