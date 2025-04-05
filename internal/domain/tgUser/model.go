package tgUser

import "gopkg.in/telebot.v4"

type TgUser struct {
	TgID         int64
	Username     string
	FirstName    string
	LastName     string
	LanguageCode string
	IsBot        bool
	IsPremium    bool
}

func ToModel(u *telebot.User) TgUser {
	return TgUser{
		TgID:         u.ID,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		LanguageCode: u.LanguageCode,
		IsBot:        u.IsBot,
		IsPremium:    u.IsPremium,
	}
}
