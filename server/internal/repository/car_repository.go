package repository

import (
	"context"
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAvailableCars(f models.CarFilter) ([]models.Car, error)
	GetCarByID(ctx context.Context, id uint) (*models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetAvailableCars(f models.CarFilter) ([]models.Car, error) {
	var cars []models.Car

	db := r.db.Where("status = ?", models.StatusAvailable)
	if f.MinPrice > 0 {
		db = db.Where("daily_price >= ?", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		db = db.Where("daily_price <= ?", f.MaxPrice)
	}
	if f.CarBrand != "" {
		db = db.Where("brand = ?", f.CarBrand)
	}
	if f.YearOfIssues > 0 {
		db = db.Where("year_of_issue = ?", f.YearOfIssues)
	}
	if len(f.CarClass) > 0 {
		db = db.Where("car_class IN ?", f.CarClass)
	}
	if err := db.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *carRepository) GetCarByID(ctx context.Context, id uint) (*models.Car, error) {
	var car models.Car
	if err := r.db.
		WithContext(ctx).
		Where("car_id = ?", id).
		First(&car).
		Error; err != nil {
		return nil, err
	}
	return &car, nil
}
