package repository

import (
	"github.com/exideys/car_rental_service/internal/models"
	"gorm.io/gorm"
)

type CarRepository interface {
	GetAvailableCars() ([]models.Car, error)
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
