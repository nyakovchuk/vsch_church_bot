package user

import (
	"context"
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
)

const (
	TelegramUsersTable = "telegram_users"
	UsersTable         = "users"
)

type Repository interface {
	RegisterUser(context.Context, tgUser.DtoRepository) error
	UpdateUserRadius(ctx context.Context, tgUserID int64, radius int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) RegisterUser(ctx context.Context, dtoTgUser tgUser.DtoRepository) error {

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

	tgUserID, err := r.createTelegramUser(ctx, tx, dtoTgUser)
	if err != nil {
		return apperrors.Wrap(apperrors.ErrInsertTelegramUser, err)
	}

	if err := r.createUser(ctx, tx, tgUserID); err != nil {
		return apperrors.Wrap(apperrors.ErrInsertUser, err)
	}

	if err := tx.Commit(); err != nil {
		return apperrors.Wrap(apperrors.ErrCommitTransaction, err)
	}
	committed = true

	return nil
}

func (r *repository) UpdateUserRadius(ctx context.Context, tgUserID int64, radius int) error {

	ds := goqu.Update(UsersTable).
		Set(goqu.Record{"radius": radius}).
		Where(goqu.Ex{"tg_user_id": tgUserID})

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

func (r *repository) createTelegramUser(ctx context.Context, tx *sql.Tx, dtoTgUser tgUser.DtoRepository) (int64, error) {
	ds := goqu.Insert(TelegramUsersTable).
		Rows(goqu.Record{
			"tg_id":         dtoTgUser.TgID,
			"username":      dtoTgUser.Username,
			"first_name":    dtoTgUser.FirstName,
			"last_name":     dtoTgUser.LastName,
			"language_code": dtoTgUser.LanguageCode,
			"is_bot":        dtoTgUser.IsBot,
			"is_premium":    dtoTgUser.IsPremium,
		}).
		Returning("tg_id")

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return 0, apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	var id int64
	err = tx.QueryRowContext(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return 0, apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return id, nil
}

func (r *repository) createUser(ctx context.Context, tx *sql.Tx, tgUserID int64) error {
	ds := goqu.Insert(UsersTable).
		Rows(goqu.Record{
			"tg_user_id": tgUserID,
		})

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return apperrors.Wrap(apperrors.ErrBuildSQL, err)
	}

	_, err = tx.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return apperrors.Wrap(apperrors.ErrExecuteQuery, err)
	}

	return nil
}
