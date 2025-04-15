package platform

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Service interface {
	GetByName(context.Context, string) (Platform, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetByName(ctx context.Context, name string) (Platform, error) {
	dtoPlatform, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return Platform{}, apperrors.Wrap(apperrors.ErrPlatformGet, err)
	}

	platform := dtoPlatform.ToModel()
	return platform, nil
}
