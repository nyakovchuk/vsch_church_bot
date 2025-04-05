package user

import (
	"context"

	"github.com/nyakovchuk/vsch_church_bot/internal/domain/tgUser"
)

type Service interface {
	CreateUser(context.Context, tgUser.TgUser) bool
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) CreateUser(ctx context.Context, modelTgUser tgUser.TgUser) bool {

	repoTgUserDto := tgUser.ModelToDto(modelTgUser)

	if err := s.repo.CreateUser(ctx, repoTgUserDto); err != nil {
		return false
	}

	return true
}
