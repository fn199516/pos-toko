package routes

import (
	"pos-toko/internal/domain"
	"pos-toko/internal/handler"
	"pos-toko/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, hlp domain.Helper) {

	hanlder := handler.NewHandler(&hlp)

	r.GET("/test", hanlder.Register)

	api := r.Group("/api/v1")
	{
		api.POST("/register", hanlder.Register)
		api.POST("/login", hanlder.Login)
	}

	middleware := middleware.NewMiddleware(&hlp)

	apiProduct := r.Group("/api/v1/product", middleware.Auth())
	{
		apiProduct.POST("", hanlder.CreateProduct)
		apiProduct.GET("", hanlder.GetProduct)
		apiProduct.DELETE("/:id", hanlder.DeleteProduct)
	}

	apiTrx := r.Group("/api/v1/transaction", middleware.Auth())
	{
		apiTrx.POST("", hanlder.CreateTransaction)
		apiTrx.GET("", hanlder.GetTransaction)
	}
}
