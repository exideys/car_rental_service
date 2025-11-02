package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestOrderHandler_Create_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockOrderService := new(mocks.MockOrderService)
	orderHandler := NewOrderHandler(mockOrderService)

	orderRequest := gin.H{
		"email":       "test@example.com",
		"car_id":      1,
		"start_date":  time.Now().Add(24 * time.Hour).Format("2006-01-02"),
		"end_date":    time.Now().Add(48 * time.Hour).Format("2006-01-02"),
		"daily_price": 100,
	}

	expectedOrder := &models.Order{
		ClientID: 1,
		CarID:    1,
	}

	mockOrderService.On("GetByEmail", "test@example.com").Return(&models.Client{ClientID: 1}, nil)
	mockOrderService.On("Create", uint(1), uint(1), uint(100), mock.Anything, mock.Anything).Return(expectedOrder, nil)

	rr := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(rr)

	jsonBody, _ := json.Marshal(orderRequest)
	req, _ := http.NewRequest(http.MethodPost, "/order", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req

	r.POST("/order", orderHandler.Create)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockOrderService.AssertExpectations(t)
}
