package rental

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RentalHandler struct {
	service *RentalService
}

func NewRentalHandler(service *RentalService) *RentalHandler {
	return &RentalHandler{service: service}
}

func (handler *RentalHandler) PostRentalHandler(context *gin.Context) {
	var newRental Rental
	var err error

	if err = context.ShouldBindJSON(&newRental); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateOneRental(&newRental); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newRental})
}

func (handler *RentalHandler) GetRentalsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	rentals, totalRecords, err := handler.service.ReadAllRentals(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": rentals, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *RentalHandler) GetRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := handler.service.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": rental})
}

func (handler *RentalHandler) PutRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := handler.service.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedRental Rental
	if err = context.ShouldBindJSON(&updatedRental); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental data format"})
		return
	}

	updatedRental.RentalID = rental.RentalID

	if err = handler.service.UpdateOneRental(&updatedRental); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rental"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedRental})
}

func (handler *RentalHandler) DeleteRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := handler.service.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Rental not found"})
		return
	}

	if err = handler.service.DeleteOneRental(rental); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete rental"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": rental})
}
