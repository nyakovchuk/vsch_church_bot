package platform

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

const (
	PlatformsTable = "platforms"
)

type Repository interface {
	GetByName(context.Context, string) (DtoRepository, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetByName(ctx context.Context, name string) (DtoRepository, error) {
	query, args, err := goqu.From(PlatformsTable).
		Select("id", "name").
		Where(goqu.C("name").Eq(name)).
		ToSQL()
	if err != nil {
		return DtoRepository{}, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	var dtoPlatform DtoRepository
	err = r.db.QueryRowContext(ctx, query, args...).
		Scan(&dtoPlatform.ID, &dtoPlatform.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return DtoRepository{}, nil
		}
		return DtoRepository{}, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return dtoPlatform, nil
}
