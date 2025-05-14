package service

import (
	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/repository"
)

type CarService interface {
	ListAvailableCars() ([]models.Car, error)
	GetFilteredCars(filters models.CarFilters) ([]models.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) ListAvailableCars() ([]models.Car, error) {
	return s.repo.GetAvailableCars()
}

func (s *carService) GetFilteredCars(filters models.CarFilters) ([]models.Car, error) {
	return s.repo.FindWithFilters(filters)
}
