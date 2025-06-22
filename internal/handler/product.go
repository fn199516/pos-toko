package handler

import (
	"encoding/json"
	"net/http"
	"pos-toko/internal/domain"
	"pos-toko/pkg/databases"

	"github.com/gin-gonic/gin"
)

func (cfg *Base) CreateProduct(c *gin.Context) {
	var (
		req     domain.ReqProduct
		product domain.Product
	)

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{StatusCode: http.StatusBadRequest, ResMessage: "INVALID_VALIDATION"})
		return
	}

	byteData, _ := json.Marshal(req)

	json.Unmarshal(byteData, &product)

	databases.DB.Create(&product)

	c.JSON(http.StatusOK, domain.Response{StatusCode: http.StatusOK, ResMessage: "SUCCESS", Data: product})

}

func (cfg *Base) GetProduct(c *gin.Context) {

	var product []domain.Product

	if err := databases.DB.Find(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, domain.Response{StatusCode: http.StatusNotFound, ResMessage: "NOT_FOUND"})
		return
	}

	c.JSON(http.StatusOK, domain.Response{StatusCode: http.StatusOK, ResMessage: "SUCCESS", Data: product})

}

func (cfg *Base) DeleteProduct(c *gin.Context) {

	id := c.Param("id")

	var product domain.Product

	if err := databases.DB.Where("id = ? ", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, domain.Response{StatusCode: http.StatusNotFound, ResMessage: "NOT_FOUND"})
		return
	}
	databases.DB.Delete(&product)

	c.JSON(http.StatusOK, domain.Response{StatusCode: http.StatusOK, ResMessage: "SUCCESS", Data: gin.H{"message": "data product berhasil di hapus", "product": product}})

}
