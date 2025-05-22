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
	c.Redirect(http.StatusSeeOther, "/authorisation")

}

func (h *AuthHandler) Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	client, err := h.svc.Login(c.Request.Context(), email, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("session_user", client.Email, 3600, "/", "", false, true)

	c.Redirect(http.StatusSeeOther, "/profile")
}

func (h *AuthHandler) GetCurrentUser(c *gin.Context) {
	email, err := c.Cookie("session_user")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authenticated"})
		return
	}

	client, err := h.svc.GetByEmail(email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"first_name": client.FirstName,
		"last_name":  client.LastName,
		"email":      client.Email,
	})
}
