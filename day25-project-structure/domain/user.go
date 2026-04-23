package domain

import "gorm.io/gorm"

// Entity - the core data structure
type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{ID: u.ID, Name: u.Name, Email: u.Email, Role: u.Role}
}

// Repository interface - defines what data operations are available
// The service depends on this interface, not on GORM directly
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	FindAll() ([]User, error)
}

// Service interface - defines what business operations are available
// The handler depends on this interface, not on the service directly
type UserService interface {
	Register(name, email, password string) (*UserResponse, error)
	Login(email, password string) (string, error)
	GetProfile(userID uint) (*UserResponse, error)
	UpdateProfile(userID uint, name string) (*UserResponse, error)
	GetAllUsers() ([]UserResponse, error)
}
