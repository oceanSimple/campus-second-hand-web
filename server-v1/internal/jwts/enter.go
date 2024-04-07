package jwts

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

// JwtPayLoad jwt中payload数据
type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"` // 用户名
	Role     int    `json:"role"`     // 权限  1 普通用户  2 管理员
}

type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}

// GetToken 创建 Token
func GetToken(user JwtPayLoad, accessSecret string, expires int64) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: user,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(accessSecret))
}

// ParseToken 解析 token
//func ParseToken(tokenStr string, accessSecret string, expires int64) (*CustomClaims, error) {
//
//	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(accessSecret), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
//		return claims, nil
//	}
//	return nil, errors.New("invalid token")
//}
