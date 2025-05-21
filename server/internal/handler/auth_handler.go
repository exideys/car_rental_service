package handler

import (
	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandler struct {
	svc service.AuthService
}

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	if err := h.svc.SignUp(
		c.Request.Context(),
		c.PostForm("first_name"),
		c.PostForm("last_name"),
		c.PostForm("email"),
		c.PostForm("telephone"),
		c.PostForm("password"),
		c.PostForm("password_confirm"),
		c.PostForm("birth_date"),
	); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/profile")
}
