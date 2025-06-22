package token

import (
	"pos-toko/internal/domain"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(jwtKey string, claims *domain.JwtClaims) (string, error) {

	exp := time.Now().Add(24 * time.Hour)

	claims.StandardClaims = jwt.StandardClaims{ExpiresAt: exp.Unix()}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtKey))

}
