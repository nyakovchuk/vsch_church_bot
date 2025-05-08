package country

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

const (
	ChurchesTable = "churches"
)

type Repository interface {
	ListCountriesByChurchesCount(context.Context) ([]WithChurchesCountDTO, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) ListCountriesByChurchesCount(ctx context.Context) ([]WithChurchesCountDTO, error) {
	query, args, err := goqu.From(ChurchesTable).
		Join(
			goqu.T("countries"),
			goqu.On(goqu.Ex{"churches.country_id": goqu.I("countries.id")}),
		).
		Select(
			goqu.I("churches.country_id"),
			goqu.I("countries.country_ru"),
			goqu.I("countries.country_en"),
			goqu.COUNT("*").As("country_count"),
		).
		GroupBy(goqu.I("churches.country_id")).
		Order(goqu.I("country_count").Desc()).
		ToSQL()
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}
	defer rows.Close()

	var countries []WithChurchesCountDTO
	for rows.Next() {
		var dto WithChurchesCountDTO
		err := rows.Scan(
			&dto.ID,
			&dto.NameRu,
			&dto.NameEn,
			&dto.ChurchesCount,
		)
		if err != nil {
			return nil, apperrors.Wrap(apperrors.ErrRowsScan, err)
		}
		countries = append(countries, dto)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Wrap(apperrors.ErrRows, err)
	}

	return countries, nil
}
