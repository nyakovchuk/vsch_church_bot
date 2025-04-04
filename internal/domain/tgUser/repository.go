package tguser

import "database/sql"

type Repository interface {
	CheckTgId(id int64) (bool, error)
}

type repository struct {
	db *sql.DB
}

func New(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CheckTgId(id int64) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE tgid = ? LIMIT 1)",
		id,
	).Scan(&exists)
	return exists, err
}
