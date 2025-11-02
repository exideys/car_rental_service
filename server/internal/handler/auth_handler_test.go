package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/exideys/car_rental_service/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/exideys/car_rental_service/internal/service/mocks"
)

func TestAuthHandler_Login_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockAuthService := new(mocks.MockAuthService)
	authHandler := NewAuthHandler(mockAuthService)

	formData := "email=test@example.com&password=password"

	mockAuthService.On("Login", mock.Anything, "test@example.com", "password").Return(&models.Client{Email: "test@example.com"}, nil)

	router := gin.New()
	router.POST("/login", authHandler.Login)

	rr := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/login", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusSeeOther, rr.Code)
	mockAuthService.AssertExpectations(t)
}
