package handler

import (
	"github.com/exideys/car_rental_service/internal/models"
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
	var f models.Car_filter
	if err := c.BindQuery(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid filter params"})
		return
	}
	cars, err := h.svc.ListAvailableCars(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}
