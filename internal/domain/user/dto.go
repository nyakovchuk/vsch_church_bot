package user

import "time"

type DtoRepository struct {
	ID        int       `db:"id"`
	TgId      int64     `db:"tg_user_id"`
	LangId    *int      `db:"lang_id"`
	Radius    int       `db:"radius"`
	CreatedAt time.Time `db:"created_at"`
}
