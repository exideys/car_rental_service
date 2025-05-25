package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/repository"
)

type OrderService interface {
	Create(clientID, carID, DailyPrice uint, start, end time.Time) (*models.Order, error)
	GetByEmail(email string) (*models.Client, error)
	GetAllOrders(email string) ([]models.Order, error)
}
type orderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{repo: repo}
}

func (s *orderService) Create(clientID, carID, DailyPrice uint, start, end time.Time) (*models.Order, error) {
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
	fmt.Println(DailyPrice)
	if dailyPrice := DailyPrice * uint(end.Sub(start).Hours()/24); dailyPrice == 0 {
		return nil, errors.New("invalid price")
	}

	order := &models.Order{
		ClientID:  clientID,
		CarID:     carID,
		StartDate: start,
		EndDate:   end,
		Status:    "Active",
		IsPaid:    true,
		TotalCost: float64(DailyPrice * uint(end.Sub(start).Hours()/24)),
	}

	if err := s.repo.Create(order); err != nil {
		return nil, err
	}
	return order, nil
}
func (s *orderService) GetByEmail(email string) (*models.Client, error) {
	a, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func (s *orderService) GetAllOrders(email string) ([]models.Order, error) {
	return s.repo.GetAllOrders(email)
}
