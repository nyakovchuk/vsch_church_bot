package country

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

const (
	ChurchesTable = "churches"
)

type Repository interface {
	ListCountriesByChurchesCount(context.Context) ([]WithChurchesCountDTO, error)
	GetFlagsFromFile() (map[string]string, error)
}

type repository struct {
	db            *sql.DB
	filenameFlags string
}

func NewRepository(db *sql.DB, filenameFlags string) Repository {
	return &repository{
		db:            db,
		filenameFlags: filenameFlags,
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

func (r *repository) GetFlagsFromFile() (map[string]string, error) {
	type Country struct {
		Name string `json:"name"`
		Flag string `json:"flag"`
	}

	file, err := os.Open(r.filenameFlags)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrOpenFile, err)
	}
	defer file.Close()

	// Декодуємо JSON у slice структур
	var countries []Country
	if err := json.NewDecoder(file).Decode(&countries); err != nil {
		return nil, apperrors.Wrap(apperrors.ErrDecodeJSON, err)
	}

	// Створюємо map: name => flag
	flagMap := make(map[string]string)
	for _, country := range countries {
		flagMap[country.Name] = country.Flag
	}

	return flagMap, nil
}
