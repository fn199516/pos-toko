package domain

import "github.com/golang-jwt/jwt/v4"

type JwtClaims struct {
	UserId uint
	Name   string
	jwt.StandardClaims
}
