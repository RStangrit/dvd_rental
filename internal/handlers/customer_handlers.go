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

func PostCustomerHandler(context *gin.Context) {
	var newCustomer models.Customer
	var err error

	if err = context.ShouldBindJSON(&newCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = services.ValidateCustomer(&newCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateCustomer(&newCustomer); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCustomer})
}

func GetCustomersHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	customers, totalRecords, err := repositories.ReadAllCustomers(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customers, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := repositories.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

func PutCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := repositories.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCustomer models.Customer
	err = context.ShouldBindJSON(&updatedCustomer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data format"})
		return
	}

	updatedCustomer.CustomerID = int(customer.CustomerID)

	err = repositories.UpdateOneCustomer(updatedCustomer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCustomer})
}

func DeleteCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := repositories.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Customer not found"})
		return
	}

	err = repositories.DeleteOneCustomer(*customer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": customer})
}
