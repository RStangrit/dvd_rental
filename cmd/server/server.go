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
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	params := config.LoadConfig()
	server := setupServer()

	registerRoutes(server)

	port := getPort(params.Port)
	log.Printf("Server is running on port %s", port)

	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupServer() *gin.Engine {
	server := gin.Default()

	server.Use(middleware.LoggerMiddleware())

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

func registerRoutes(server *gin.Engine) {
	routes := []func(*gin.Engine){
		language.RegisterLanguageRoutes,
		actor.RegisterActorRoutes,
		film.RegisterFilmRoutes,
		category.RegisterCategoryRoutes,
		film_actor.RegisterFilmActorRoutes,
		inventory.RegisterInventoryRoutes,
		film_category.RegisterFilmCategoryRoutes,
		country.RegisterCountryRoutes,
		city.RegisterCityRoutes,
		address.RegisterAddressRoutes,
		customer.RegisterCustomerRoutes,
		staff.RegisterStaffRoutes,
		store.RegisterStoreRoutes,
		rental.RegisterRentalRoutes,
		payment.RegisterPaymentRoutes,
		user.RegisterUserRoutes,
		file.RegisterFileRoutes,
		development.RegisterDevelopmentRoutes,
	}

	for _, register := range routes {
		register(server)
	}
}
