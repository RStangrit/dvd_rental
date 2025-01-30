package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/internal/services"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostAddressHandler(context *gin.Context) {
	var newAddress models.Address
	var err error

	if err = context.ShouldBindJSON(&newAddress); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = services.ValidateAddress(&newAddress); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateAddress(&newAddress); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newAddress})
}

func GetAddressesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	addresses, totalRecords, err := repositories.ReadAllAddresses(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": addresses, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := repositories.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": address})
}

func PutAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := repositories.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedAddress models.Address
	err = context.ShouldBindJSON(&updatedAddress)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address data format"})
		return
	}

	if err = services.ValidateAddress(&updatedAddress); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAddress.AddressID = int(address.AddressID)

	err = repositories.UpdateOneAddress(updatedAddress)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedAddress})
}

func DeleteAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := repositories.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = repositories.DeleteOneAddress(*address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": address})
}
