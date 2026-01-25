package middleware

import (
	"auth-system/internal/pkg"

	"github.com/gin-gonic/gin"
)

// RequestID 请求ID中间件
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID, err := pkg.GenerateRandomString(16)
		if err != nil {
			requestID = "unknown"
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}
