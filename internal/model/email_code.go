package model

import (
	"time"
)

// EmailCodeType 邮箱验证码类型
type EmailCodeType string

const (
	EmailCodeTypeRegister EmailCodeType = "register" // 注册
	EmailCodeTypeLogin    EmailCodeType = "login"    // 登录
	EmailCodeTypeReset    EmailCodeType = "reset"    // 重置密码
)

// EmailCode 邮箱验证码
type EmailCode struct {
	ID           uint          `gorm:"primarykey" json:"id"`
	Email        string        `gorm:"index;not null;size:100" json:"email"`
	Code         string        `gorm:"not null;size:10" json:"code"`
	Type         EmailCodeType `gorm:"not null;size:20;index" json:"type"`
	IP           string        `gorm:"size:45" json:"ip"`
	Used         bool          `gorm:"not null;default:false;index" json:"used"`
	FailCount    int           `gorm:"not null;default:0" json:"fail_count"`
	ExpireAt     time.Time     `gorm:"not null;index" json:"expire_at"`
	CreatedAt    time.Time     `json:"created_at"`
}

func (EmailCode) TableName() string {
	return "email_codes"
}
