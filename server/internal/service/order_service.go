package service

import (
	"errors"
	"time"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(clientID, carID uint, start, end time.Time) (*models.Order, error) {
	if clientID == 0 {
		return nil, errors.New("unauthorized: clientID is required")
	}

	now := time.Now()
	if start.Before(now) {
		return nil, errors.New("invalid rental period: start date is in the past")
	}
	if !end.After(start) {
		return nil, errors.New("invalid rental period: end date must be after start date")
	}

	order := &models.Order{
		ClientID:  clientID,
		CarID:     carID,
		StartDate: start,
		EndDate:   end,
		Status:    "Active",
		IsPaid:    true,
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}
