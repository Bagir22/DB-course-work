package service

type Repository interface {
}

type Service struct {
	repo Repository
}

func InitService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}
