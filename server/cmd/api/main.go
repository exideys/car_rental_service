package main

import (
	"github.com/exideys/car_rental_service/configs"
	"github.com/exideys/car_rental_service/internal/handler"
	"github.com/exideys/car_rental_service/internal/repository"
	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	//Connection to a database
	db := configs.InitDB()
	//Repositories
	carRepo := repository.NewCarRepository(db)
	authRepo := repository.NewAuthRepository(db)
	//Services
	carSvc := service.NewCarService(carRepo)
	authSvc := service.NewAuthService(authRepo)
	//Handlers
	carHandler := handler.NewCarHandler(carSvc)
	authHandler := handler.NewAuthHandler(authSvc)
	//Router
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	// --- sessions ---
	store := cookie.NewStore([]byte("super-secret-key"))
	r.Use(sessions.Sessions("car_rental_session", store))
	// --- end sessions ---
	if err != nil {
		return
	}
	r.Static("/css", "../Client/css")
	r.Static("/js", "../Client/js")
	r.StaticFile("/catalog", "../Client/html/catalog.html")
	r.StaticFile("", "../Client/html/index.html")
	r.StaticFile("/contacts", "../Client/html/contacts.html")
	r.StaticFile("/about_us", "../Client/html/about.html")
	r.StaticFile("/profile", "../Client/html/profile.html")
	r.StaticFile("/authorisation", "../Client/html/login.html")
	r.POST("/authorisation", authHandler.SignUp)
	api := r.Group("/api")
	{
		api.GET("/cars", carHandler.GetAvailableCars)
	}
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
