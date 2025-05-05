package handler

import (
	"net/http"

	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-gonic/gin"
)

type CarHandler struct {
	svc service.CarService
}

func NewCarHandler(svc service.CarService) *CarHandler {
	return &CarHandler{svc: svc}
}

func (h *CarHandler) GetAvailableCars(c *gin.Context) {
	cars, err := h.svc.ListAvailableCars()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}
