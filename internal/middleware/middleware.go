package middleware

import (
	"fmt"
	"net/http"
	"pos-toko/internal/domain"
	"pos-toko/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type Middleware struct {
	Helper *domain.Helper
}

func NewMiddleware(hlp *domain.Helper) Middleware {
	return Middleware{Helper: hlp}
}

func (cfg *Middleware) Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		log := logger.Log
		tokenStr := ctx.GetHeader("Authorization")
		if tokenStr == "" {
			ctx.JSON(http.StatusUnauthorized, domain.Response{StatusCode: http.StatusUnauthorized, ResMessage: "Unauthorized"})
			ctx.Abort()
			return
		}

		claims := &domain.JwtClaims{}

		fmt.Println("token str : ", tokenStr)
		fmt.Println("secret key : ", cfg.Helper.JwtSecret)
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {

			return []byte(cfg.Helper.JwtSecret), nil

		})

		if err != nil || !token.Valid {
			log.Error("error validation token", zap.String("err", err.Error()))
			ctx.JSON(http.StatusUnauthorized, domain.Response{StatusCode: http.StatusUnauthorized, ResMessage: "Unauthorized"})
			ctx.Abort()
			return
		}

		log.Info("UserId", zap.Any("userId", claims.UserId))
		ctx.Set("UserId", claims.UserId)
		ctx.Next()

	}
}
