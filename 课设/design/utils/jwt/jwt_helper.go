package jwt

import (
	"design/config"
	"encoding/json"
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// 解码token
type DecodedToken struct {
	Iat      int    `json:"iat"` // Iat 表示该 JWT 的签发时间
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Iss      string `json:"iss"` // Iss 表示该 JWT 的签发者
	IsAdmin  bool   `json:"isAdmin"`
}

var ErrNotToken = errors.New("JWT token missing")

// GenerateToken 函数用于生成 JSON Web Token (JWT)。
func GenerateToken(claims *jwt.Token, secret string) (token string) {
	// 将 secret 转换成字节数组
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	// 通过 SignedString 方法生成签名后的 JWT
	token, _ = claims.SignedString(hmacSecret)

	return
}

// VerifyToken 函数用于验证 JWT 是否有效，并返回解码后的 Token 信息。
func VerifyToken(token string, secret string) *DecodedToken {
	// 将 secret 转换成字节数组
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	// 解析 JWT 并验证签名
	decoded, err := jwt.Parse(
		token, func(token *jwt.Token) (interface{}, error) {
			return hmacSecret, nil
		})

	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	// 将解码后的 JWT 转换成结构体 DecodedToken
	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	jsonErr := json.Unmarshal(jsonString, &decodedToken)
	if jsonErr != nil {
		log.Print(jsonErr)
	}

	return &decodedToken
}

func Decoded(tokenString string) (string, error) {
	if tokenString == "" {
		return "", ErrNotToken
	}

	// 移除"Bearer "前缀（如果存在）
	const bearerPrefix = "Bearer "
	if len(tokenString) > len(bearerPrefix) && strings.EqualFold(tokenString[:len(bearerPrefix)], bearerPrefix) {
		tokenString = tokenString[len(bearerPrefix):]
	}

	// 解析JWT令牌
	token := VerifyToken(tokenString, config.SecretKey)
	if token == nil {
		// 如果解析失败，返回错误
		return "", ErrNotToken
	}
	return token.UserId, nil
}
