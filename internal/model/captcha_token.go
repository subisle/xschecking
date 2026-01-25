package model

import (
	"time"
)

// CaptchaToken 验证码通行证（一次性凭证）
type CaptchaToken struct {
	ID         uint      `gorm:"primarykey" json:"id"`
	Token      string    `gorm:"uniqueIndex;not null;size:64" json:"token"`
	CaptchaID  string    `gorm:"not null;size:64" json:"captcha_id"`
	IP         string    `gorm:"index;size:45" json:"ip"`
	UAHash     string    `gorm:"size:64" json:"ua_hash"`
	Used       bool      `gorm:"not null;default:false;index" json:"used"`
	ExpireAt   time.Time `gorm:"not null;index" json:"expire_at"`
	CreatedAt  time.Time `json:"created_at"`
}

func (CaptchaToken) TableName() string {
	return "captcha_tokens"
}
