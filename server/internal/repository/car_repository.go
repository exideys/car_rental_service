package repository

import (
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAvailableCars() ([]models.Car, error)
	FindWithFilters(filters models.CarFilters) ([]models.Car, error)
}

type carRepository struct {
	db *gorm.DB
}

func NewCarRepository(db *gorm.DB) CarRepository {
	return &carRepository{db: db}
}

func (r *carRepository) GetAvailableCars() ([]models.Car, error) {
	var cars []models.Car
	if err := r.db.
		Where("status = ?", models.StatusAvailable).
		Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *carRepository) FindWithFilters(f models.CarFilters) ([]models.Car, error) {
	var cars []models.Car
	query := r.db.Model(&models.Car{}).Where("status = ?", models.StatusAvailable)

	if f.Brand != "" {
		query = query.Where("brand = ?", f.Brand)
	}
	if f.TransportType != "" {
		query = query.Where("transport_type = ?", f.TransportType)
	}
	if f.BodyType != "" {
		query = query.Where("body_type = ?", f.BodyType)
	}
	if f.Transmission != "" {
		query = query.Where("transmission = ?", f.Transmission)
	}
	if f.MinPrice > 0 {
		query = query.Where("daily_price >= ?", f.MinPrice)
	}
	if f.MaxPrice > 0 {
		query = query.Where("daily_price <= ?", f.MaxPrice)
	}

	if err := query.Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}
