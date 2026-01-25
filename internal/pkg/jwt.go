package pkg

import (
	"auth-system/internal/config"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims JWT声明
type Claims struct {
	UserID     uint   `json:"uid"`
	LoginAt    int64  `json:"login_at"`
	IPSegment  string `json:"ip_segment"`
	UAHash     string `json:"ua_hash"`
	jwt.RegisteredClaims
}

// GenerateToken 生成JWT Token
func GenerateToken(userID uint, ip string, userAgent string) (string, error) {
	cfg := config.GetConfig()
	
	now := time.Now()
	expiresAt := now.Add(time.Duration(cfg.JWT.ExpireHours) * time.Hour)

	claims := Claims{
		UserID:    userID,
		LoginAt:   now.Unix(),
		IPSegment: getIPSegment(ip),
		UAHash:    hashString(userAgent),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    cfg.JWT.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(cfg.JWT.Secret))
}

// ParseToken 解析JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	cfg := config.GetConfig()

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// getIPSegment 获取IP段（前三段）
func getIPSegment(ip string) string {
	// 简单实现，实际可以更复杂
	if len(ip) > 0 {
		for i := len(ip) - 1; i >= 0; i-- {
			if ip[i] == '.' {
				return ip[:i]
			}
		}
	}
	return ip
}

// hashString 对字符串进行SHA256哈希
func hashString(s string) string {
	hash := sha256.Sum256([]byte(s))
	return hex.EncodeToString(hash[:])
}
