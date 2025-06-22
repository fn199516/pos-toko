package handler

import (
	"net/http"
	"pos-toko/internal/domain"
	"pos-toko/pkg/databases"

	"github.com/gin-gonic/gin"
)

func (cfg *Base) CreateTransaction(c *gin.Context) {

	var (
		req     domain.ReqTransaction
		product domain.Product
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "INVALID_VALIDATION"})
		return
	}

	userId := c.MustGet("UserId").(uint)

	if err := databases.DB.Where("id = ? ", req.ProductId).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, domain.Response{StatusCode: http.StatusNotFound, ResMessage: "NOT_FOUND"})
		return
	}

	if product.Stock < req.Quantity {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "INVALID_REQUEST", Data: gin.H{"messgae": "stock kurang dari quantity"}})
		return
	}

	totalPrice := product.Price * float64(req.Quantity)

	product.Stock -= req.Quantity

	trx := domain.Transaction{
		UserID:     userId,
		ProductId:  req.ProductId,
		Quantity:   req.Quantity,
		TotalPrice: totalPrice,
	}

	databases.DB.Save(&product)
	databases.DB.Create(&trx)
	c.JSON(http.StatusCreated, domain.Response{StatusCode: http.StatusCreated, ResMessage: "SUCCESS", Data: trx})

}

func (cfg *Base) GetTransaction(c *gin.Context) {
	var trx []domain.Transaction

	userId := c.MustGet("UserId").(uint)
	databases.DB.Where("user_id = ? ", userId).Find(&trx)

	c.JSON(http.StatusOK, domain.Response{StatusCode: http.StatusOK, ResMessage: "SUCCESS", Data: trx})

}
