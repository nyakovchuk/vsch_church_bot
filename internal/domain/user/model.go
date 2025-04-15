package user

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
)

type User struct {
	ID           int
	Platform     platform.Platform
	ExternalId   string
	TgID         int64
	Coordinates  model.Coordinates
	LangId       *int
	Radius       int
	Username     string
	FirstName    string
	LastName     string
	LanguageCode string
	IsBot        bool
	IsPremium    bool
}
