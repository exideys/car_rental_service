package mocks

import (
	"time"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) Create(clientID, carID, DailyPrice uint, start, end time.Time) (*models.Order, error) {
	args := m.Called(clientID, carID, DailyPrice, start, end)
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) GetByEmail(email string) (*models.Client, error) {
	args := m.Called(email)
	return args.Get(0).(*models.Client), args.Error(1)
}

func (m *MockOrderService) GetAllOrders(email string) ([]models.Order, error) {
	args := m.Called(email)
	return args.Get(0).([]models.Order), args.Error(1)
}
