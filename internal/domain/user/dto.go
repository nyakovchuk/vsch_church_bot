package user

import "time"

type DtoRepository struct {
	ID           int       `db:"id"`
	PlatformId   int       `db:"platform_id"`
	ExternalId   string    `db:"external_id"`
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
