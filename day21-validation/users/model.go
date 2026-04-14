package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id"         gorm:"primaryKey;autoIncrement"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Name      string         `json:"name"`
	Email     string         `json:"email"      gorm:"uniqueIndex"`
	Age       int            `json:"age"`
}

type CreateUserInput struct {
	Name  string `json:"name"  binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age"   binding:"required,min=1,max=120"`
}

type UpdateUserInput struct {
	Name  string `json:"name"  binding:"omitempty,min=2,max=50"`
	Email string `json:"email" binding:"omitempty,email"`
	Age   int    `json:"age"   binding:"omitempty,gte=1,lte=120"`
}
