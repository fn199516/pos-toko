package databases

import (
	"log"
	"pos-toko/pkg/logger"

	mysql_driver "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection(dsn string) {

	gormDb, err := gorm.Open(mysql_driver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal Koneksi ke database server : ", err)
	}
	DB = gormDb

	log := logger.Log

	log.Info("success connection to database")
}
