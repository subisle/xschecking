package middleware

import (
	"auth-system/internal/pkg"
	"strings"

	"github.com/gin-gonic/gin"
)

// Auth JWT认证中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			pkg.Error(c, pkg.CodeUnauthorized, "未提供认证令牌")
			c.Abort()
			return
		}

		// Bearer token
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			pkg.Error(c, pkg.CodeUnauthorized, "认证令牌格式错误")
			c.Abort()
			return
		}

		// 解析token
		claims, err := pkg.ParseToken(parts[1])
		if err != nil {
			pkg.Error(c, pkg.CodeTokenInvalid, "无效的认证令牌")
			c.Abort()
			return
		}

		// 验证IP和UA（可选，根据安全需求）
		currentIP := c.ClientIP()
		currentUA := c.Request.UserAgent()
		
		// 简单的IP段验证
		if !strings.HasPrefix(currentIP, claims.IPSegment) {
			pkg.Error(c, pkg.CodeUnauthorized, "IP地址变化，请重新登录")
			c.Abort()
			return
		}

		// 将用户信息存入上下文
		c.Set("user_id", claims.UserID)
		c.Set("login_at", claims.LoginAt)
		c.Set("ua_hash", currentUA)
		
		c.Next()
	}
}
