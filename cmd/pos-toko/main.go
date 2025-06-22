package main

import (
	"fmt"
	"log"
	"pos-toko/configs"
	"pos-toko/internal/domain"
	"pos-toko/internal/routes"
	"pos-toko/pkg/databases"
	"pos-toko/pkg/logger"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("pos-toko")
	var cfg domain.Helper

	configs.LoadEnv()
	configs.SetHelper(&cfg)

	logger.InitLoggerZap()

	// koneksi database
	databases.Connection(cfg.DbDsn)

	databases.DB.AutoMigrate(&domain.Product{}, &domain.User{}, &domain.Transaction{})
	r := gin.Default()
	routes.Register(r, cfg)

	if err := r.Run(":" + cfg.Port); err != nil {
		log.Fatal("error run apps pos-toko")
	}

}
