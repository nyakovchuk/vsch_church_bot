package language

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

const (
	LanguagesTable = "languages"
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
	query, args, err := goqu.From(LanguagesTable).
		Select("id", "name", "code").
		ToSQL()
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}
	defer rows.Close()

	var languages []DtoRepository
	for rows.Next() {
		var dto DtoRepository
		err := rows.Scan(
			&dto.ID,
			&dto.Name,
			&dto.Code,
		)
		if err != nil {
			return nil, apperrors.Wrap(apperrors.ErrRowsScan, err)
		}
		languages = append(languages, dto)
	}

	if err = rows.Err(); err != nil {
		return nil, apperrors.Wrap(apperrors.ErrRows, err)
	}

	return languages, nil
}
