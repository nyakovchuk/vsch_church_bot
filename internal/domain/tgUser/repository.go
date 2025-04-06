package tgUser

import (
	"database/sql"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Repository interface {
	CheckTgId(id int64) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CheckTgId(id int64) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM telegram_users WHERE tg_id = ? LIMIT 1)",
		id,
	).Scan(&exists)

	if err != nil {
		return false, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return exists, nil
}
