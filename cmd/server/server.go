package server

import (
	"log"
	"main/config"
	"main/internal/actor"
	"main/internal/address"
	"main/internal/category"
	"main/internal/city"
	"main/internal/country"
	"main/internal/customer"
	"main/internal/development"
	"main/internal/file"
	"main/internal/film"
	"main/internal/film_actor"
	"main/internal/film_category"
	"main/internal/inventory"
	"main/internal/language"
	"main/internal/payment"
	"main/internal/rental"
	"main/internal/staff"
	"main/internal/store"
	user "main/internal/user"
	"main/middleware"
	"main/pkg/websocket"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitServer(db *gorm.DB) {
	params := config.LoadConfig()
	server := setupServer()

	registerRoutes(server, db)

	port := getPort(params.Port)
	log.Printf("Server is running on port %s", port)

	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupServer() *gin.Engine {
	server := gin.Default()

	server.Use(middleware.TimeTrackerMiddleware(), middleware.CorsMiddleware(), middleware.LoggerMiddleware())

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	return server
}

func getPort(configPort string) string {
	if configPort != "" {
		return configPort
	}
	if envPort := os.Getenv("PORT"); envPort != "" {
		return envPort
	}
	return "8080"
}

func registerRoutes(server *gin.Engine, db *gorm.DB) {
	//registration method for MVC
	addressRoutes := address.NewAddressRoutes(db)
	addressRoutes.RegisterAddressRoutes(server)
	actorRoutes := actor.NewActorRoutes(db)
	actorRoutes.RegisterActorRoutes(server)
	categoryRoutes := category.NewCategoryRoutes(db)
	categoryRoutes.RegisterCategoryRoutes(server)
	cityRoutes := city.NewCityRoutes(db)
	cityRoutes.RegisterCityRoutes(server)

	routes := []func(*gin.Engine){

		//old registration method
		language.RegisterLanguageRoutes,
		film.RegisterFilmRoutes,
		film_actor.RegisterFilmActorRoutes,
		inventory.RegisterInventoryRoutes,
		film_category.RegisterFilmCategoryRoutes,
		country.RegisterCountryRoutes,
		customer.RegisterCustomerRoutes,
		staff.RegisterStaffRoutes,
		store.RegisterStoreRoutes,
		rental.RegisterRentalRoutes,
		payment.RegisterPaymentRoutes,
		user.RegisterUserRoutes,
		file.RegisterFileRoutes,
		development.RegisterDevelopmentRoutes,
		websocket.RegisterWSRoutes,
	}

	for _, register := range routes {
		register(server)
	}
}
