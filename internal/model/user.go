package model

import (
	"time"
)

// UserStatus 用户状态
type UserStatus int

const (
	UserStatusNormal  UserStatus = 1 // 正常
	UserStatusFrozen  UserStatus = 2 // 冻结
	UserStatusBanned  UserStatus = 3 // 封禁
)

// User 用户模型
type User struct {
	ID           uint       `gorm:"primarykey" json:"id"`
	Email        string     `gorm:"uniqueIndex;not null;size:100" json:"email"`
	PasswordHash string     `gorm:"not null;size:255" json:"-"`
	Status       UserStatus `gorm:"not null;default:1" json:"status"`
	LastLoginAt  *time.Time `gorm:"index" json:"last_login_at"`
	LastLoginIP  string     `gorm:"size:45" json:"last_login_ip"`
	CreatedAt    time.Time  `gorm:"index" json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

func (User) TableName() string {
	return "users"
}
