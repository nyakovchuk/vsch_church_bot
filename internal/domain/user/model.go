package user

import (
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/model"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/platform"
	"github.com/nyakovchuk/vsch_church_bot/utils"
	"gopkg.in/telebot.v4"
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

func FromTelebotUser(u *telebot.User) User {
	tgIdText := utils.Int64ToString(u.ID)

	return User{
		ExternalId:   tgIdText,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		LanguageCode: u.LanguageCode,
		IsBot:        u.IsBot,
		IsPremium:    u.IsPremium,
	}
}

func (u *User) SetPlatform(platform platform.Platform) {
	u.Platform = platform
}
