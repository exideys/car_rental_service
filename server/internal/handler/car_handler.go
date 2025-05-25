package handler

import (
	"github.com/exideys/car_rental_service/internal/models"
	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type request struct {
	ID uint `form:"car_id" binding:"required"`
}
type CarHandler struct {
	svc service.CarService
}

func NewCarHandler(svc service.CarService) *CarHandler {
	return &CarHandler{svc: svc}
}

func (h *CarHandler) GetAvailableCars(c *gin.Context) {
	var f models.CarFilter
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

func (h *CarHandler) GetCar(c *gin.Context) {
	var req request
	if err := c.BindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid car id"})
		return
	}

	car, err := h.svc.GetCar(c.Request.Context(), req.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "car not found"})
		return
	}
	c.JSON(http.StatusOK, car)
}
