package handler

import (
	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type OrderHandler interface {
	Create(c *gin.Context)
	GetAllOrders(c *gin.Context)
}
type OrderRequest struct {
	Email     string `json:"email"`
	CarID     uint   `json:"car_id" binding:"required"`
	StartDate string `json:"start_date" binding:"required,datetime=2006-01-02"`
	EndDate   string `json:"end_date" binding:"required,datetime=2006-01-02"`
}

type orderHandler struct {
	svc service.OrderService
}

func NewOrderHandler(svc service.OrderService) OrderHandler {
	return &orderHandler{svc: svc}
}

func (h *orderHandler) Create(c *gin.Context) {
	/*coockie, err := c.Cookie("current_user")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err.Error()"})
		return
	}
	decodedValue, err := url.QueryUnescape(coockie)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var data map[string]string
	if err := json.Unmarshal([]byte(decodedValue), &data); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	email, ok := data["email"]
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized: invalid email in context"})
		return
	}*/
	//const payload = {
	//      email:    'sad@gmail.com',
	//      car_id:     Number(fd.get('car_id')),
	//      start_date: fd.get('start_date'),
	//      end_date:   fd.get('end_date')
	//    };

	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}
	client, err := h.svc.GetByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "client not found"})
		return
	}

	start, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format"})
		return
	}
	end, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format"})
		return
	}

	order, err := h.svc.Create(client.ClientID, req.CarID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func (h *orderHandler) GetAllOrders(c *gin.Context) {
	var email string
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	orderss, err := h.svc.GetAllOrders(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderss)
}
