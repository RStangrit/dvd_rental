package address

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	service *AddressService
}

func NewAddressHandler(service *AddressService) *AddressHandler {
	return &AddressHandler{service: service}
}

func (handler *AddressHandler) PostAddressHandler(context *gin.Context) {
	var newAddress Address
	var err error

	if err = context.ShouldBindJSON(&newAddress); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateAddress(&newAddress); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newAddress})
}

func (handler *AddressHandler) GetAddressesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	addresses, totalRecords, err := handler.service.ReadAllAddresses(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": addresses, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *AddressHandler) GetAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := handler.service.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": address})
}

func (handler *AddressHandler) PutAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := handler.service.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedAddress Address
	err = context.ShouldBindJSON(&updatedAddress)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address data format"})
		return
	}

	updatedAddress.AddressID = int(address.AddressID)

	err = handler.service.UpdateOneAddress(&updatedAddress)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update address"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedAddress})
}

func (handler *AddressHandler) DeleteAddressHandler(context *gin.Context) {
	addressId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address ID format"})
		return
	}

	address, err := handler.service.ReadOneAddress(addressId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneAddress(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete address"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": address})
}
