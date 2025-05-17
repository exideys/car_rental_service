package service

import (
	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/repository"
)

type CarService interface {
	ListAvailableCars(f models.CarFilter) ([]models.Car, error)
}

type carService struct {
	repo repository.CarRepository
}

func NewCarService(repo repository.CarRepository) CarService {
	return &carService{repo: repo}
}

func (s *carService) ListAvailableCars(f models.CarFilter) ([]models.Car, error) {
	return s.repo.GetAvailableCars(f)
}
