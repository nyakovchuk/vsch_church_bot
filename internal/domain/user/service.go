package user

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) CheckTgId(tgId int) bool {
	return true
}
