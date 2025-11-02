package mocks

import (
	"context"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockCarService struct {
	mock.Mock
}

func (m *MockCarService) ListAvailableCars(f models.CarFilter) ([]models.Car, error) {
	args := m.Called(f)
	return args.Get(0).([]models.Car), args.Error(1)
}

func (m *MockCarService) GetCar(ctx context.Context, id uint) (*models.Car, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*models.Car), args.Error(1)
}
