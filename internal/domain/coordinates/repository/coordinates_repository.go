package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

const CoordinatesTable = "coordinates"

type CoordinatesRepository interface {
	Save(context.Context, *dto.RepositoryCoordinates) (*dto.RepositoryCoordinates, error)
	GetByID(ctx context.Context, id int) (*dto.RepositoryCoordinates, error)
}

type coordinatesRepository struct {
	db *sql.DB
}

func NewCoordinatesRepository(db *sql.DB) *coordinatesRepository {
	return &coordinatesRepository{
		db: db,
	}
}

func (r *coordinatesRepository) Save(ctx context.Context, coords *dto.RepositoryCoordinates) (*dto.RepositoryCoordinates, error) {
	coords.CreatedAt = time.Now().UTC()

	ds := goqu.Insert(CoordinatesTable).
		Rows(goqu.Record{
			"latitude":   coords.Latitude,
			"longitude":  coords.Longitude,
			"is_on_text": coords.IsOnText,
			"created_at": coords.CreatedAt,
		}).
		Returning("id")

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	// сохранить в users coordinates_id

	// Написать ещё два запроса для таблицы user, telegram и coordinates_history,

	// Выполняем через стандартный sql.DB
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBeginTransaction, err)
	}
	defer tx.Rollback()

	var id int64
	err = tx.QueryRowContext(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	if err := tx.Commit(); err != nil {
		return nil, apperrors.Wrap(apperrors.ErrCommitTransaction, err)
	}

	coords.ID = int(id)
	return coords, nil
}

func (r *coordinatesRepository) GetByID(ctx context.Context, id int) (*dto.RepositoryCoordinates, error) {
	ds := goqu.From(CoordinatesTable).
		Select("id", "latitude", "longitude", "created_at").
		Where(goqu.C("id").Eq(id)).
		Limit(1)

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	var result dto.RepositoryCoordinates
	var createdAt string

	row := r.db.QueryRowContext(ctx, sqlQuery, args...)
	err = row.Scan(
		&result.ID,
		&result.Latitude,
		&result.Longitude,
		&createdAt,
	)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	result.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrParseTime, err)
	}

	return &result, nil
}
