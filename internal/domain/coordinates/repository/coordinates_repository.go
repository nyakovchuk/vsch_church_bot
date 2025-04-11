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
	Save(context.Context, *dto.RepositoryCoordinates) (dto.RepositoryCoordinates, error)
	GetUserId(ctx context.Context, tgUserId int64) (dto.RepositoryCoordinates, error)
}

type coordinatesRepository struct {
	db *sql.DB
}

func NewCoordinatesRepository(db *sql.DB) *coordinatesRepository {
	return &coordinatesRepository{
		db: db,
	}
}

func (r *coordinatesRepository) Save(ctx context.Context, coords *dto.RepositoryCoordinates) (dto.RepositoryCoordinates, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrBeginTransaction, err)
	}

	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	// если нет координат в user_id - создать
	// если есть - обновить

	savedCoords, err := r.createOrUpdateCoordinates(ctx, tx, coords)
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrInsertCoordinates, err)
	}

	// сохранить в users coordinates_id

	// сохранить координаты в coordinates_history,

	if err := tx.Commit(); err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrCommitTransaction, err)
	}
	committed = true

	return savedCoords, nil

}

func (r *coordinatesRepository) GetUserId(ctx context.Context, tgUserId int64) (dto.RepositoryCoordinates, error) {
	ds := goqu.From(CoordinatesTable).
		Select("latitude", "longitude").
		Where(goqu.C("tg_user_id").Eq(tgUserId)).
		Limit(1)

	query, args, err := ds.Prepared(true).ToSQL()
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	var latitude, longitude float64
	err = r.db.QueryRowContext(ctx, query, args...).Scan(&latitude, &longitude)
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	var result dto.RepositoryCoordinates
	result.Latitude = latitude
	result.Longitude = longitude

	return result, nil
}

func (r *coordinatesRepository) createOrUpdateCoordinates(ctx context.Context, tx *sql.Tx, coords *dto.RepositoryCoordinates) (dto.RepositoryCoordinates, error) {
	ds := goqu.Insert(CoordinatesTable).
		Rows(goqu.Record{
			"tg_user_id": coords.TgUserID,
			"latitude":   coords.Latitude,
			"longitude":  coords.Longitude,
			"is_on_text": coords.IsOnText,
		}).
		OnConflict(goqu.DoUpdate("tg_user_id", goqu.Record{
			"latitude":   coords.Latitude,
			"longitude":  coords.Longitude,
			"is_on_text": coords.IsOnText,
			"updated_at": time.Now().Format("2006-01-02 15:04:05"),
		})).
		Returning("id", "tg_user_id", "latitude", "longitude")

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	var result dto.RepositoryCoordinates
	row := tx.QueryRowContext(ctx, sqlQuery, args...)
	err = row.Scan(
		&result.ID,
		&result.TgUserID,
		&result.Latitude,
		&result.Longitude,
	)
	if err != nil {
		return dto.RepositoryCoordinates{}, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return result, nil
}
