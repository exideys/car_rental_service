package repository

import (
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAvailableCars(f models.Car_filter) ([]models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetAvailableCars(f models.Car_filter) ([]models.Car, error) {
	var cars []models.Car

	db := r.db.Where("status = ?", models.StatusAvailable)
	if f.Min_price > 0 {
		db = db.Where("daily_price >= ?", f.Min_price)
	}
	if f.Max_price > 0 {
		db = db.Where("daily_price <= ?", f.Max_price)
	}
	if f.Car_brand != "" {
		db = db.Where("brand = ?", f.Car_brand)
	}
	if f.Year_of_issues > 0 {
		db = db.Where("year_of_issue = ?", f.Year_of_issues)
	}
	if len(f.Car_class) > 0 {
		db = db.Where("car_class IN ?", f.Car_class)
	}
	if err := db.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}
