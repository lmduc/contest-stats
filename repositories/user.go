package repositories

import (
	"time"
)

// User holds the attributes of user
type User struct {
	ID             *string    `gorm:"column:id"`
	CreatedAt      *time.Time `gorm:"column:created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at"`
	LeetcodeHandle *string    `gorm:"column:leetcode_handle"`
}

// NewUser generates a new User with prefilled CreatedAt
func NewUser() *User {
	now := time.Now()
	return &User{
		CreatedAt: &now,
	}
}
