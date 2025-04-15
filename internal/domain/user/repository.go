package user

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/external"
)

const (
	UsersTable = "users"
)

type Repository interface {
	IsRegistered(platformId int, externalId string) (bool, error)
	RegisterUser(context.Context, DtoRepository) error
	UpdateUserRadius(ctx context.Context, external external.ExternalRepository, radius int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) IsRegistered(platformId int, externalId string) (bool, error) {
	var exists bool
	err := r.db.QueryRow(
		"SELECT EXISTS(SELECT 1 FROM users WHERE platform_id = ? AND external_id = ? LIMIT 1)",
		platformId, externalId,
	).Scan(&exists)

	if err != nil {
		return false, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return exists, nil
}

func (r *repository) RegisterUser(ctx context.Context, dtoUser DtoRepository) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return apperrors.Wrap(apperrors.ErrBeginTransaction, err)
	}

	// Отложенный Rollback с проверкой
	var committed bool
	defer func() {
		if !committed {
			tx.Rollback()
		}
	}()

	if err := r.createUser(ctx, tx, dtoUser); err != nil {
		return apperrors.Wrap(apperrors.ErrInsertUser, err)
	}

	if err := tx.Commit(); err != nil {
		return apperrors.Wrap(apperrors.ErrCommitTransaction, err)
	}
	committed = true

	return nil
}

func (r *repository) UpdateUserRadius(ctx context.Context, external external.ExternalRepository, radius int) error {

	ds := goqu.Update(UsersTable).
		Set(goqu.Record{"radius": radius}).
		Where(goqu.Ex{"platform_id": external.PlatformID, "external_id": external.ID})

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	_, err = r.db.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return nil
}

func (r *repository) createUser(ctx context.Context, tx *sql.Tx, dtoUser DtoRepository) error {
	sqlQuery, args, err := goqu.
		Insert(UsersTable).
		Rows(goqu.Record{
			"platform_id":   dtoUser.PlatformID,
			"external_id":   dtoUser.ExternalID,
			"username":      dtoUser.Username,
			"first_name":    dtoUser.FirstName,
			"last_name":     dtoUser.LastName,
			"language_code": dtoUser.LanguageCode,
			"is_bot":        dtoUser.IsBot,
			"is_premium":    dtoUser.IsPremium,
		}).
		ToSQL()
	if err != nil {
		return apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	_, err = tx.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return nil
}
