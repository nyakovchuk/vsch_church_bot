package user

import "time"

type DtoRepository struct {
	ID           int       `db:"id"`
	PlatformID   int       `db:"platform_id"`
	ExternalID   string    `db:"external_id"`
	LangId       *int      `db:"lang_id"`
	Radius       int       `db:"radius"`
	Username     string    `db:"username"`
	FirstName    string    `db:"first_name"`
	LastName     string    `db:"last_name"`
	LanguageCode string    `db:"language_code"`
	IsBot        bool      `db:"is_bot"`
	IsPremium    bool      `db:"is_premium"`
	CreatedAt    time.Time `db:"created_at"`
}

func ToDto(u User) DtoRepository {
	return DtoRepository{
		PlatformID:   u.Platform.ID,
		ExternalID:   u.ExternalId,
		LangId:       u.LangId,
		Radius:       u.Radius,
		Username:     u.Username,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		LanguageCode: u.LanguageCode,
		IsBot:        u.IsBot,
		IsPremium:    u.IsPremium,
	}
}
