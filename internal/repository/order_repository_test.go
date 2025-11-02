package repository

import (
	"testing"
	"time"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestOrderRepository_Create(t *testing.T) {
	order := &models.Order{
		ID:        1,
		CarID:     1,
		StartDate: time.Now(),
		EndDate:   time.Now().Add(24 * time.Hour),
		Status:    "active",
	}

	// ...existing code...
}