package tgUser

type Service interface {
	CheckTgId(tgId int64) bool
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) CheckTgId(tgId int64) bool {
	exist, err := s.repo.CheckTgId(tgId)
	if err != nil || !exist {
		return false
	}
	return true
}
