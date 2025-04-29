package language

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Service interface {
	GetAll(context.Context) ([]Language, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]Language, error) {
	dtoLanguages, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrLanguageGetAll, err)
	}

	languages := ToModels(&dtoLanguages)
	return languages, err
}
