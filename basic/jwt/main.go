/*
@see https://www.liwenzhou.com/posts/Go/jwt_in_gin/
1. 依赖：jwt-go
*/
package main

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// TokenExpireDuration 过期时间
const TokenExpireDuration = time.Hour * 2

// MySecret 密钥
var MySecret = []byte("hello")

// MyClaims 载荷 定义jwt中要保存的信息
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken 生成JWT
func GenToken(username string) (string, error) {
	// 创建一个我们自己的声明
	c := MyClaims{
		username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "golang",
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

func main() {
	token, err := GenToken("root")
	if err != nil {
		fmt.Println("cannot generate jwt, err:", err)
	}
	fmt.Println("jwt:", token)

	claims, err := ParseToken(token)
	if err != nil {
		fmt.Println("cannot parse jwt, err:", err)
	}
	fmt.Println(claims.Username, claims.ExpiresAt, claims.Issuer)
}
