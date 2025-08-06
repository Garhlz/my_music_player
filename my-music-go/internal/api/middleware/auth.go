package middleware

import (
	"my-music-go/internal/config" // 1. 引入 config 包
	"my-music-go/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 2. 删除硬编码的常量
// const JWT_SECRET = "" 

// NewAuthMiddleware 创建一个 Gin 中间件用于 JWT 认证
// 3. 这是一个工厂函数，它接收 config 作为参数
func NewAuthMiddleware(cfg config.Config) gin.HandlerFunc {
	// 4. 它返回一个闭包，这个内部函数就是我们真正的中间件
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "请求未包含访问令牌"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌格式无效"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, services.ErrInvalidCredentials
			}
			// 5. 使用从外部传入的 cfg.JWTSecret
			return []byte(cfg.JWTSecret), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌无效或已过期"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, ok := claims["user_id"].(float64); ok {
				c.Set("userID", int64(userID))
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "令牌中的用户ID无效"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "无法解析令牌声明"})
			c.Abort()
			return
		}

		c.Next()
	}
}