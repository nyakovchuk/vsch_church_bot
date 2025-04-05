package user

import "time"

type DtoRepository struct {
	ID            int       `db:"id"`
	TgId          int64     `db:"telegram_users_id"`
	CoordinatesId *int      `db:"coordinates_id"`
	LangId        *int      `db:"lang_id"`
	Radius        int       `db:"radius"`
	CreatedAt     time.Time `db:"created_at"`
}
