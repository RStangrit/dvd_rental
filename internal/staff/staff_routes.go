package staff

import (
	"github.com/gin-gonic/gin"
)

func RegisterStaffRoutes(server *gin.Engine) {
	server.POST("/staff", PostStaffHandler)
	server.GET("/staffs", GetStaffsHandler)
	server.GET("/staff/:id", GetStaffHandler)
	server.PUT("/staff/:id", PutStaffHandler)
	server.DELETE("/staff/:id", DeleteStaffHandler)
}
