package tgUser

type Service interface {
	CheckTgId(tgId int64) (bool, error)
}

type service struct {
	repo Repository
}

func NewService(repository Repository) Service {
	return &service{
		repo: repository,
	}
}

func (s *service) CheckTgId(tgId int64) (bool, error) {
	exist, err := s.repo.CheckTgId(tgId)
	if err != nil {
		return false, err
	}

	if !exist {
		return false, nil
	}

	return true, nil
}
