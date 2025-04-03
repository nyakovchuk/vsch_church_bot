package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/coordinates/dto"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
)

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

	// Строим запрос с помощью goqu
	ds := goqu.Insert("coordinates").
		Rows(goqu.Record{
			"latitude":   coords.Latitude,
			"longitude":  coords.Longitude,
			"created_at": coords.CreatedAt,
		}).
		Returning("id")

	// Получаем SQL и аргументы
	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
	}

	// Выполняем через стандартный sql.DB
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var id int64
	err = tx.QueryRowContext(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	coords.ID = int(id)
	return coords, nil
}

// Пример метода для получения данных
func (r *coordinatesRepository) GetByID(ctx context.Context, id int) (*dto.RepositoryCoordinates, error) {
	ds := goqu.From("coordinates").
		Select("id", "latitude", "longitude", "created_at").
		Where(goqu.C("id").Eq(id)).
		Limit(1)

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	// Парсим время из SQLite
	result.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAt)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
