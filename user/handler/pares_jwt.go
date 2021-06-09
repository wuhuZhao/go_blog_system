package user_handler

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"hkzhao/go_blog_system/config"
	user "hkzhao/go_blog_system/user/struct"
	"time"
)

func ToJwt(username string) (string, error) {
	maxAge := 60 * 60 * 24
	userClaims := &user.UserClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(maxAge) * time.Second).Unix(),
			Issuer:    username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)
	tokenString, err := token.SignedString([]byte(config.SECRETKEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func ParseToken(tokenString string) (*user.UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &user.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected string method: %v ", token.Header["alg"])
		}
		return []byte(config.SECRETKEY), nil
	})
	if claims, ok := token.Claims.(*user.UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
