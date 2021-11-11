package tool

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// MySecret 密钥
var MySecret = []byte("hello")

// MyClaims 载荷 定义jwt中要保存的信息
type MyClaims struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`

	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(claims MyClaims) (string, error) {
	// 载荷
	c := MyClaims{
		ID:       claims.ID,
		Username: claims.Username,
		Nickname: claims.Nickname,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "kratos",
		},
	}

	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(MySecret)
}

// ParseToken 解析token
func ParseToken(tokenStr string) (*MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
