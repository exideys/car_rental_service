package configs

import (
	"log"

	"github.com/exideys/car_rental_service/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/rent_cars?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	if err := db.AutoMigrate(&models.Car{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
	return db
}
