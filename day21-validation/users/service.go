package users

import (
	"errors"

	"gorm.io/gorm"
)

var ErrNotFound = errors.New("user not found")
var ErrDeleted = errors.New("user is deleted, cannot update")

type Service struct {
	DB *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{DB: db}
}

func (s *Service) GetAll() []User {
	var users []User
	s.DB.Unscoped().Find(&users)
	return users
}

func (s *Service) Create(input CreateUserInput) (User, error) {
	user := User{Name: input.Name, Age: input.Age, Email: input.Email}
	result := s.DB.Create(&user)
	return user, result.Error
}

func (s *Service) Delete(id string) error {
	var user User
	if result := s.DB.First(&user, id); result.Error != nil {
		return ErrNotFound
	}
	s.DB.Unscoped().Delete(&user)
	return nil
}

func (s *Service) PurgeSoftDeleted() int64 {
	result := s.DB.Unscoped().Where("deleted_at IS NOT NULL").Delete(&User{})
	return result.RowsAffected
}

func (s *Service) Update(id string, input UpdateUserInput) (User, error) {
	var user User
	if result := s.DB.Unscoped().First(&user, id); result.Error != nil {
		return user, ErrNotFound
	}
	if user.DeletedAt.Valid {
		return user, ErrDeleted
	}
	s.DB.Model(&user).Updates(input)
	s.DB.First(&user, user.ID)
	return user, nil
}
