package staff

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type StaffRoutes struct {
	handler *StaffHandler
}

func NewStaffRoutes(db *gorm.DB) *StaffRoutes {
	repo := NewStaffRepository(db)
	service := NewStaffService(repo)
	handler := NewStaffHandler(service)

	return &StaffRoutes{handler: handler}
}

func (route *StaffRoutes) RegisterStaffRoutes(server *gin.Engine) {
	server.POST("/staff", route.handler.PostStaffHandler)
	server.GET("/staffs", route.handler.GetStaffsHandler)
	server.GET("/staff/:id", route.handler.GetStaffHandler)
	server.PUT("/staff/:id", route.handler.PutStaffHandler)
	server.DELETE("/staff/:id", route.handler.DeleteStaffHandler)
}
