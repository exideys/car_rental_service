package handler

import (
	"net/http"
	"strconv"

	"github.com/exideys/car_rental_service/internal/models"
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
	filters := models.CarFilters{
		Brand:         c.Query("brand"),
		TransportType: c.Query("transport_type"),
		BodyType:      c.Query("body_type"),
		Transmission:  c.Query("transmission"),
		MinPrice:      parseFloat(c.Query("min_price")),
		MaxPrice:      parseFloat(c.Query("max_price")),
	}

	cars, err := h.svc.GetFilteredCars(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func parseFloat(value string) float64 {
	f, _ := strconv.ParseFloat(value, 64)
	return f
}
