package service

import (
	"errors"

	"go-basics/day25-project-structure/domain"
	"go-basics/day25-project-structure/utils"

	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(name, email, password string) (*domain.UserResponse, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("could not hash password")
	}

	user := &domain.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
		Role:     "user",
	}

	if err := s.repo.Create(user); err != nil {
		return nil, errors.New("email already registered")
	}

	res := user.ToResponse()
	return &res, nil
}

func (s *userService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}

func (s *userService) GetProfile(userID uint) (*domain.UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	res := user.ToResponse()
	return &res, nil
}

func (s *userService) UpdateProfile(userID uint, name string) (*domain.UserResponse, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	user.Name = name
	if err := s.repo.Update(user); err != nil {
		return nil, errors.New("could not update user")
	}
	res := user.ToResponse()
	return &res, nil
}

func (s *userService) GetAllUsers() ([]domain.UserResponse, error) {
	users, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}
	var result []domain.UserResponse
	for _, u := range users {
		result = append(result, u.ToResponse())
	}
	return result, nil
}
