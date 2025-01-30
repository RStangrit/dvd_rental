package routes

import (
	"main/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterStaffRoutes(server *gin.Engine) {
	server.POST("/staff", handlers.PostStaffHandler)
	server.GET("/staffs", handlers.GetStaffsHandler)
	server.GET("/staff/:id", handlers.GetStaffHandler)
	server.PUT("/staff/:id", handlers.PutStaffHandler)
	server.DELETE("/staff/:id", handlers.DeleteStaffHandler)
}
