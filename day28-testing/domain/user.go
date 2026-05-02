package domain

// Plain struct — no GORM dependency, easy to mock and test
type User struct {
	ID       uint
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

// UserRepository — the interface the service depends on.
// Swap the real GORM impl for a mock in tests.
type UserRepository interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
	FindByID(id uint) (*User, error)
	Update(user *User) error
	FindAll() ([]User, error)
}

type UserService interface {
	Register(name, email, password string) (*UserResponse, error)
	Login(email, password string) (string, error)
	GetProfile(userID uint) (*UserResponse, error)
	UpdateProfile(userID uint, name string) (*UserResponse, error)
	GetAllUsers() ([]UserResponse, error)
}
