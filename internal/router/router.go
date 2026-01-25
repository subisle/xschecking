package router

import (
	"auth-system/internal/handler"
	"auth-system/internal/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(mode string) *gin.Engine {
	// 设置运行模式
	gin.SetMode(mode)

	r := gin.New()

	// 全局中间件
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", handler.HealthCheck)

		// 验证码相关
		captcha := v1.Group("/captcha")
		{
			captcha.POST("/generate", handler.GenerateCaptcha)
			captcha.POST("/verify", handler.VerifyCaptcha)
		}

		// 认证相关（不需要JWT）
		auth := v1.Group("/auth")
		{
			auth.POST("/register", handler.Register)
			auth.POST("/login", handler.Login)
		}

		// 邮箱相关
		email := v1.Group("/email")
		{
			email.POST("/send-code", handler.SendEmailCode)
		}

		// 用户相关（需要JWT认证）
		user := v1.Group("/user")
		user.Use(middleware.Auth())
		{
			user.GET("/info", handler.GetUserInfo)
			user.PUT("/password", handler.UpdatePassword)
		}

		// 管理后台（需要JWT认证和管理员权限）
		admin := v1.Group("/admin")
		admin.Use(middleware.Auth())
		{
			admin.GET("/users", handler.ListUsers)
			admin.PUT("/users/:id/status", handler.UpdateUserStatus)
		}
	}

	return r
}
