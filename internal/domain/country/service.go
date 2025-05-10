package country

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Service interface {
	ListCountriesByChurchesCount(context.Context) ([]CountryWithChurchesCount, error)
	FetchCountryChurchesStats(context.Context) ([]CountryWithChurchesCount, error)
	GetFlags() (map[string]string, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) FetchCountryChurchesStats(ctx context.Context) ([]CountryWithChurchesCount, error) {

	countries, err := s.ListCountriesByChurchesCount(ctx)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrListCountriesByChurchesCount, err)
	}

	// це не критична помилка, тому повертаємо список країн
	// але без прапорів
	mapFlags, err := s.GetFlags()
	if err != nil {
		return countries, apperrors.Wrap(apperrors.ErrGetFlags, err)
	}

	for i := range countries {
		countryName := countries[i].NameEn
		countries[i].Flag = mapFlags[countryName]
	}

	return countries, nil
}

func (s *service) ListCountriesByChurchesCount(ctx context.Context) ([]CountryWithChurchesCount, error) {
	dtoCountriesWithChurches, err := s.repo.ListCountriesByChurchesCount(ctx)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrLanguageGetAll, err)
	}

	countries := ToModels(&dtoCountriesWithChurches)
	return countries, nil
}

func (s *service) GetFlags() (map[string]string, error) {
	return s.repo.GetFlagsFromFile()
}
