package utils

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateToken 生成 JWT
func GenerateToken(userID int64, secretKey string) (string, error) {
	// 创建一个新的 token 对象
	token := jwt.New(jwt.SigningMethodHS256)

	// 设置 claims
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 令牌有效期 24 小时
	claims["iat"] = time.Now().Unix()                      // 签发时间

	// 使用密钥签名并获取完整的编码后的字符串 token
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}