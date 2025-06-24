package database

import (
	"fmt"
	"log"
	"user-service/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/user_service?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal konek DB:", err)
	}
	DB.AutoMigrate(&models.User{})
	fmt.Println("Database connected.")
}
