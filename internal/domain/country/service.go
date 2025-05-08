package country

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Service interface {
	ListCountriesByChurchesCount(context.Context) ([]CountryWithChurchesCount, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) ListCountriesByChurchesCount(ctx context.Context) ([]CountryWithChurchesCount, error) {
	dtoCountriesWithChurches, err := s.repo.ListCountriesByChurchesCount(ctx)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrLanguageGetAll, err)
	}

	countries := ToModels(&dtoCountriesWithChurches)
	return countries, nil
}
