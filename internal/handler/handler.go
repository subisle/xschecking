package handler

import (
	"auth-system/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "服务运行正常",
	})
}

// GenerateCaptcha 生成验证码
func GenerateCaptcha(c *gin.Context) {
	// TODO: 实现验证码生成逻辑
	pkg.Success(c, gin.H{
		"captcha_id": "placeholder",
		"captcha_url": "placeholder",
	})
}

// VerifyCaptcha 验证验证码
func VerifyCaptcha(c *gin.Context) {
	// TODO: 实现验证码校验逻辑
	pkg.Success(c, gin.H{
		"captcha_pass_token": "placeholder",
	})
}

// Register 用户注册
func Register(c *gin.Context) {
	// TODO: 实现用户注册逻辑
	// 1. 校验 captcha_pass_token
	// 2. 校验邮箱验证码
	// 3. 校验卡密（如果需要）
	// 4. 风控判断
	// 5. 创建用户
	pkg.Success(c, gin.H{
		"message": "注册成功",
	})
}

// Login 用户登录
func Login(c *gin.Context) {
	// TODO: 实现用户登录逻辑
	// 1. 校验 captcha_pass_token
	// 2. 验证用户名密码
	// 3. 风控检查
	// 4. 生成JWT
	pkg.Success(c, gin.H{
		"token": "placeholder",
	})
}

// SendEmailCode 发送邮箱验证码
func SendEmailCode(c *gin.Context) {
	// TODO: 实现邮箱验证码发送逻辑
	pkg.Success(c, gin.H{
		"message": "验证码已发送",
	})
}

// GetUserInfo 获取用户信息
func GetUserInfo(c *gin.Context) {
	// TODO: 实现获取用户信息逻辑
	userID := c.GetUint("user_id")
	pkg.Success(c, gin.H{
		"user_id": userID,
		"email": "placeholder@example.com",
	})
}

// UpdatePassword 修改密码
func UpdatePassword(c *gin.Context) {
	// TODO: 实现修改密码逻辑
	pkg.Success(c, gin.H{
		"message": "密码修改成功",
	})
}

// ListUsers 获取用户列表（管理员）
func ListUsers(c *gin.Context) {
	// TODO: 实现获取用户列表逻辑
	pkg.Success(c, gin.H{
		"users": []interface{}{},
		"total": 0,
	})
}

// UpdateUserStatus 更新用户状态（管理员）
func UpdateUserStatus(c *gin.Context) {
	// TODO: 实现更新用户状态逻辑
	pkg.Success(c, gin.H{
		"message": "用户状态更新成功",
	})
}
