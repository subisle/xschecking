package pkg

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 响应码定义
const (
	CodeSuccess            = 0
	CodeError              = 1
	CodeInvalidParams      = 400
	CodeUnauthorized       = 401
	CodeForbidden          = 403
	CodeNotFound           = 404
	CodeTooManyRequests    = 429
	CodeInternalError      = 500
	CodeCaptchaError       = 1001
	CodeEmailCodeError     = 1002
	CodeUserExists         = 1003
	CodeUserNotFound       = 1004
	CodePasswordError      = 1005
	CodeAccountFrozen      = 1006
	CodeAccountBanned      = 1007
	CodeTokenExpired       = 1008
	CodeTokenInvalid       = 1009
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
	})
}

// ErrorWithData 带数据的错误响应
func ErrorWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
