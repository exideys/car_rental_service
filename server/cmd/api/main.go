package main

import (
	"github.com/exideys/car_rental_service/configs"
	"github.com/exideys/car_rental_service/internal/handler"
	"github.com/exideys/car_rental_service/internal/repository"
	"github.com/exideys/car_rental_service/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()
	carRepo := repository.NewCarRepository(db)
	carSvc := service.NewCarService(carRepo)
	carHandler := handler.NewCarHandler(carSvc)
	r := gin.Default()
	err := r.SetTrustedProxies(nil)
	if err != nil {
		return
	}
	r.Static("/css", "../Client/css")
	r.Static("/js", "../Client/js")
	r.StaticFile("/catalog", "../Client/html/catalog.html")
	api := r.Group("/api")
	{
		api.GET("/cars", carHandler.GetAvailableCars)
	}
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
