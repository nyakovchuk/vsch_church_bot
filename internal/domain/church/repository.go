package church

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
	GetAll(context.Context) ([]DtoRepository, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]DtoRepository, error) {
	ds := goqu.From(ChurchesTable).
		Select(
			"churches.id",
			"churches.name_ru",
			"churches.name_en",
			"churches.alias",
			"churches.country_ru",
			"churches.country_id",
			"churches.state_id",
			"churches.city_id",
			"churches.address_ru",
			"churches.latitude",
			"churches.longitude",
			"churches.confession_id",
			goqu.I("conf.name_ru").As("confession_name"),
		).
		Join(
			goqu.T("confessions").As("conf"),
			goqu.On(goqu.I("churches.confession_id").Eq(goqu.I("conf.id"))),
		).
		Where(goqu.I("churches.is_hidden").Eq(0))

	query, args, err := ds.ToSQL()
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}
	defer rows.Close()

	var churches []DtoRepository
	for rows.Next() {
		var dto DtoRepository
		err := rows.Scan(
			&dto.ID,
			&dto.NameRu,
			&dto.NameEn,
			&dto.Alias,
			&dto.CountryRu,
			&dto.CountryId,
			&dto.StateId,
			&dto.CityId,
			&dto.AddressRu,
			&dto.Latitude,
			&dto.Longitude,
			&dto.ConfessionId,
			&dto.ConfessionName,
		)
		if err != nil {
			return nil, apperrors.Wrap(apperrors.ErrRowsScan, err)
		}
		churches = append(churches, dto)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Wrap(apperrors.ErrRows, err)
	}

	return churches, nil
}
