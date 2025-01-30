package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostRentalHandler(context *gin.Context) {
	var newRental models.Rental
	var err error

	if err = context.ShouldBindJSON(&newRental); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateRental(&newRental); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newRental})
}

func GetRentalsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	rentals, totalRecords, err := repositories.ReadAllRentals(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": rentals, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := repositories.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": rental})
}

func PutRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := repositories.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedRental models.Rental
	err = context.ShouldBindJSON(&updatedRental)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental data format"})
		return
	}

	updatedRental.RentalID = rental.RentalID

	err = repositories.UpdateOneRental(updatedRental)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update rental"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedRental})
}

func DeleteRentalHandler(context *gin.Context) {
	rentalID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid rental ID format"})
		return
	}

	rental, err := repositories.ReadOneRental(rentalID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Rental not found"})
		return
	}

	err = repositories.DeleteOneRental(*rental)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete rental"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": rental})
}
