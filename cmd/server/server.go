package server

import (
	"context"
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

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitServer(db *gorm.DB, redisClient *redis.Client) {
	params := config.LoadConfig()
	server := setupServer(redisClient)

	registerRoutes(server, db)

	port := getPort(params.Port)
	log.Printf("Server is running on port %s", port)

	if err := server.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

func setupServer(redisClient *redis.Client) *gin.Engine {
	server := gin.Default()

	server.Use(middleware.TimeTrackerMiddleware(), middleware.CorsMiddleware(), middleware.LoggerMiddleware())

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	server.GET("/set/:key/:value", func(c *gin.Context) {
		key := c.Param("key")
		value := c.Param("value")
		ctx := context.Background()

		err := redisClient.Set(ctx, key, value, 0).Err()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Value set", "key": key, "value": value})
	})

	server.GET("/get/:key", func(c *gin.Context) {
		key := c.Param("key")
		ctx := context.Background()

		value, err := redisClient.Get(ctx, key).Result()
		if err == redis.Nil {
			c.JSON(404, gin.H{"error": "Key not found"})
			return
		} else if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"key": key, "value": value})
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
	addressRoutes := address.NewAddressRoutes(db)
	addressRoutes.RegisterAddressRoutes(server)
	actorRoutes := actor.NewActorRoutes(db)
	actorRoutes.RegisterActorRoutes(server)
	categoryRoutes := category.NewCategoryRoutes(db)
	categoryRoutes.RegisterCategoryRoutes(server)
	cityRoutes := city.NewCityRoutes(db)
	cityRoutes.RegisterCityRoutes(server)
	filmActorRoutes := film_actor.NewFilmActorRoutes(db)
	filmActorRoutes.RegisterFilmActorRoutes(server)
	inventoryRoutes := inventory.NewInventoryRoutes(db)
	inventoryRoutes.RegisterInventoryRoutes(server)
	languageRoutes := language.NewLanguageRoutes(db)
	languageRoutes.RegisterLanguageRoutes(server)
	filmRoutes := film.NewFilmRoutes(db)
	filmRoutes.RegisterFilmRoutes(server)
	countryRoutes := country.NewCountryRoutes(db)
	countryRoutes.RegisterCountryRoutes(server)
	filmCategoryRoutes := film_category.NewFilmCategoryRoutes(db)
	filmCategoryRoutes.RegisterFilmCategoryRoutes(server)
	customerRoutes := customer.NewCustomerRoutes(db)
	customerRoutes.RegisterCustomerRoutes(server)
	staffRoutes := staff.NewStaffRoutes(db)
	staffRoutes.RegisterStaffRoutes(server)
	storeRoutes := store.NewStoreRoutes(db)
	storeRoutes.RegisterStoreRoutes(server)
	rentalRoutes := rental.NewRentalRoutes(db)
	rentalRoutes.RegisterRentalRoutes(server)
	paymentRoutes := payment.NewPaymentRoutes(db)
	paymentRoutes.RegisterPaymentRoutes(server)
	userRoutes := user.NewUserRoutes(db)
	userRoutes.RegisterUserRoutes(server)
	fileRoutes := file.NewFileRoutes(db)
	fileRoutes.RegisterFileRoutes(server)
	developmentRoutes := development.NewDevelopmentRoutes(db)
	developmentRoutes.RegisterDevelopmentRoutes(server)
	websocketRoutes, _ := websocket.NewWebSocketRoutes()
	websocketRoutes.RegisterWSRoutes(server)
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
