package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"     gorm:"not null"`
	Email    string `json:"email"    gorm:"uniqueIndex;not null"`
	Password string `json:"-"        gorm:"not null"`
	Role     string `json:"role"     gorm:"default:user"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
}
