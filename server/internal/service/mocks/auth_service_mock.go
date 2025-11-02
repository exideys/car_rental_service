package mocks

import (
	"context"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/stretchr/testify/mock"
)

type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) SignUp(ctx context.Context, firstName, lastName, email, telephone, password, passwordConfirm, birthDate string) error {
	args := m.Called(ctx, firstName, lastName, email, telephone, password, passwordConfirm, birthDate)
	return args.Error(0)
}

func (m *MockAuthService) Login(ctx context.Context, email, password string) (*models.Client, error) {
	args := m.Called(ctx, email, password)
	return args.Get(0).(*models.Client), args.Error(1)
}

func (m *MockAuthService) GetByEmail(email string) (*models.Client, error) {
	args := m.Called(email)
	return args.Get(0).(*models.Client), args.Error(1)
}
