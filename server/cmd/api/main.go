package main

import (
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/modules/config"
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	_ "github.com/GoAdminGroup/go-admin/template/chartjs"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/exideys/car_rental_service/configs"
	"github.com/exideys/car_rental_service/internal/admin/tables"
	_ "github.com/exideys/car_rental_service/internal/admin/tables"
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
	orderRepo := repository.NewOrderRepository(db)
	//Services
	carSvc := service.NewCarService(carRepo)
	authSvc := service.NewAuthService(authRepo)
	orderSvc := service.NewOrderService(orderRepo)
	//Handlers
	carHandler := handler.NewCarHandler(carSvc)
	authHandler := handler.NewAuthHandler(authSvc)
	orderHandler := handler.NewOrderHandler(orderSvc)
	//Router
	r := gin.Default()
	err := r.SetTrustedProxies(nil)

	eng := engine.Default()
	// --- sessions ---
	store := cookie.NewStore([]byte("super-secret-key"))
	r.Use(sessions.Sessions("car_rental_session", store))
	// --- end sessions ---
	if err != nil {
		return
	}
	r.Static("/css", "../Client/css")
	r.Static("/js", "../Client/js")
	r.Static("/assets", "../Client/assets")
	r.StaticFile("/catalog", "../Client/html/catalog.html")
	r.StaticFile("", "../Client/html/index.html")
	r.StaticFile("/contacts", "../Client/html/contacts.html")
	r.StaticFile("/about_us", "../Client/html/about.html")
	r.StaticFile("/profile", "../Client/html/profile.html")
	r.StaticFile("/authorisation", "../Client/html/login.html")
	r.Static("/html", "../Client/html")
	r.POST("/authorisation", authHandler.SignUp)
	r.POST("/order", orderHandler.Create)
	r.POST("/login", authHandler.Login)
	r.POST("/logout", func(c *gin.Context) {
		c.SetCookie("session_user", "", -1, "/", "", false, true)
		c.Status(200)
	})
	api := r.Group("/api")
	{
		api.GET("/cars", carHandler.GetAvailableCars)
		api.GET("/current_user", authHandler.GetCurrentUser)
	}
	cfg := &config.Config{
		Databases: config.DatabaseList{
			"default": {
				Host:   "localhost",
				Port:   "3306",
				User:   "root",
				Pwd:    "",
				Name:   "rent_cars",
				Driver: "mysql",
			},
		},
		UrlPrefix: "admin",
		Theme:     "adminlte",
		IndexUrl:  "/info/cars",
		Store: config.Store{
			Path:   "./uploads",
			Prefix: "uploads",
		},
	}
	adminPlugin := admin.NewAdmin()

	for name, gen := range tables.Generators {
		adminPlugin.AddGenerator(name, gen)
	}

	if err := eng.
		AddConfig(cfg).
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}

	err = r.Run(":8080")
	if err != nil {
		return
	}
}
