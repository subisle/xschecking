package middleware

import (
	"auth-system/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Recovery 崩溃恢复中间件
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				pkg.GetLogger().Error("Panic recovered",
					zap.String("request_id", c.GetString("request_id")),
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
				)

				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    pkg.CodeInternalError,
					"message": "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
