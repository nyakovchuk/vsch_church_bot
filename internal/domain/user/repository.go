package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
)

type Repository interface {
	CreateUser(context.Context, tgUser.DtoRepository) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateUser(ctx context.Context, dtoTgUser tgUser.DtoRepository) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
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
		return fmt.Errorf("failed to insert telegram user: %w", err)
	}

	if err := r.createUser(ctx, tx, tgUserID); err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	committed = true

	return nil
}

func (r *repository) createTelegramUser(ctx context.Context, tx *sql.Tx, dtoTgUser tgUser.DtoRepository) (int64, error) {
	ds := goqu.Insert("telegram_users").
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
		return 0, fmt.Errorf("failed to build SQL: %w", err)
	}

	var id int64
	err = tx.QueryRowContext(ctx, sqlQuery, args...).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return id, nil
}

func (r *repository) createUser(ctx context.Context, tx *sql.Tx, tgUserID int64) error {
	ds := goqu.Insert("users").
		Rows(goqu.Record{
			"telegram_users_id": tgUserID,
			"created_at":        time.Now().UTC(),
		})

	sqlQuery, args, err := ds.ToSQL()
	if err != nil {
		return fmt.Errorf("failed to build SQL: %w", err)
	}

	_, err = tx.ExecContext(ctx, sqlQuery, args...)
	if err != nil {
		return fmt.Errorf("failed to execute query: %w", err)
	}

	return nil
}
