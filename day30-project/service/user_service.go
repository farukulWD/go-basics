package service

import (
	"day30-project/domain"
	"day30-project/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req domain.CreateUserRequest) (*domain.User, error) {
	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id uint) (*domain.User, error) {
	return s.repo.FindByID(id)
}
