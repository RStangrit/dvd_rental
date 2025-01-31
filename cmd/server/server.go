package server

import (
	"main/internal/actor"
	"main/internal/address"
	"main/internal/category"
	"main/internal/city"
	"main/internal/country"
	"main/internal/customer"
	"main/internal/film"
	"main/internal/film_actor"
	"main/internal/film_category"
	"main/internal/inventory"
	"main/internal/language"
	"main/internal/payment"
	"main/internal/rental"
	"main/internal/staff"
	"main/internal/store"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() {
	server := gin.Default()

	server.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	language.RegisterLanguageRoutes(server)
	actor.RegisterActorRoutes(server)
	film.RegisterFilmRoutes(server)
	category.RegisterCategoryRoutes(server)
	film_actor.RegisterFilmActorRoutes(server)
	inventory.RegisterInventoryRoutes(server)
	film_category.RegisterFilmCategoryRoutes(server)
	country.RegisterCountryRoutes(server)
	city.RegisterCityRoutes(server)
	address.RegisterAddressRoutes(server)
	customer.RegisterCustomerRoutes(server)
	staff.RegisterStaffRoutes(server)
	store.RegisterStoreRoutes(server)
	rental.RegisterRentalRoutes(server)
	payment.RegisterPaymentRoutes(server)
	server.Run()
}
