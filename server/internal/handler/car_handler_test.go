package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCarHandler_GetAvailableCars_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCarService := new(mocks.MockCarService)
	carHandler := NewCarHandler(mockCarService)

	expectedCars := []models.Car{
		{Brand: "Audi", Model: "A4"},
	}

	mockCarService.On("ListAvailableCars", mock.Anything).Return(expectedCars, nil)

	rr := httptest.NewRecorder()
	ctx, r := gin.CreateTestContext(rr)

	req, _ := http.NewRequest(http.MethodGet, "/api/cars", nil)
	ctx.Request = req

	r.GET("/api/cars", carHandler.GetAvailableCars)
	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var actualCars []models.Car
	json.Unmarshal(rr.Body.Bytes(), &actualCars)
	assert.Equal(t, expectedCars, actualCars)

	mockCarService.AssertExpectations(t)
}
