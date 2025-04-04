package tguser

type Service interface {
	CheckTgId(tgId int) bool
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) CheckTgId(tgId int) bool {
	return true
}
