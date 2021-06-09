package user

import "github.com/dgrijalva/jwt-go"

type Admin struct {
	Username string `gorm:"primaryKey"`
	Password string
	Auth     string
}

type UserClaims struct {
	Username string
	jwt.StandardClaims
}
