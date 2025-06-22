package handler

import (
	"encoding/json"
	"net/http"
	"pos-toko/internal/domain"
	"pos-toko/pkg/databases"
	"pos-toko/pkg/logger"
	"pos-toko/pkg/token"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (cfg *Base) Login(c *gin.Context) {

	var (
		req  domain.ReqLogin
		user domain.User
	)

	log := logger.Log

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "INVALID_VALIDATION"})
		return
	}

	if err := databases.DB.Where("username = ? or email = ? ", req.Username, req.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, domain.Response{StatusCode: http.StatusNotFound, ResMessage: "NOT_FOUND"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Info("compare password not match")
		c.JSON(http.StatusNotFound, domain.Response{StatusCode: http.StatusNotFound, ResMessage: "NOT_FOUND"})
		return
	}

	token, err := token.GenerateToken(cfg.Helper.JwtSecret, &domain.JwtClaims{UserId: user.ID, Name: user.Name})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "errorMessage": err.Error()})
		return
	}
	c.JSON(http.StatusOK, domain.Response{StatusCode: http.StatusOK, ResMessage: "SUCCESS", Data: gin.H{"token": token}})

}

func (cfg *Base) Register(c *gin.Context) {

	var (
		req  domain.ReqRegisterUser
		user domain.User
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "INVALID_VALIDATION"})
		return
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	byteReq, _ := json.Marshal(req)
	json.Unmarshal(byteReq, &user)

	user.Password = string(hashPassword)

	if err := databases.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "ERROR_CREATED_ACCOUNT"})

		return
	}

	c.JSON(http.StatusCreated, domain.Response{StatusCode: http.StatusCreated, ResMessage: "SUCCESS", Data: user})

}
