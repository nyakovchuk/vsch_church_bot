package church

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/apperrors"
)

type Service interface {
	GetAll(context.Context) ([]Church, error)
	// GetById(id int) (Church, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}

func (s *service) GetAll(ctx context.Context) ([]Church, error) {
	dtoChurches, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, apperrors.Wrap(apperrors.ErrChurchGetAll, err)
	}

	churches := ToModels(&dtoChurches)
	return churches, nil
}
