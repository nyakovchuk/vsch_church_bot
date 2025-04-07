package user

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
)

type Service interface {
	Register(context.Context, tgUser.TgUser) error
	UpdateUserRadius(ctx context.Context, tgUserID int64, radius int) error
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) Register(ctx context.Context, modelTgUser tgUser.TgUser) error {

	repoTgUserDto := tgUser.ModelToDto(modelTgUser)

	if err := s.repo.RegisterUser(ctx, repoTgUserDto); err != nil {
		return apperrors.Wrap(apperrors.ErrUserRegistration, err)
	}

	return nil
}

func (s *service) UpdateUserRadius(ctx context.Context, tgUserID int64, radius int) error {

	if err := s.repo.UpdateUserRadius(ctx, tgUserID, radius); err != nil {
		return apperrors.Wrap(apperrors.ErrUpdateRadius, err)
	}

	return nil
}
