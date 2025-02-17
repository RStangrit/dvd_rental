package customer

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCustomerHandler(context *gin.Context) {
	var newCustomer Customer
	var err error

	if err = context.ShouldBindJSON(&newCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateCustomer(&newCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateCustomer(db.GORM, &newCustomer); err != nil {
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

	customers, totalRecords, err := ReadAllCustomers(db.GORM, pagination)
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

	customer, err := ReadOneCustomer(db.GORM, customerId)
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

	customer, err := ReadOneCustomer(db.GORM, customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCustomer Customer
	err = context.ShouldBindJSON(&updatedCustomer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data format"})
		return
	}

	updatedCustomer.CustomerID = int(customer.CustomerID)

	err = UpdateOneCustomer(db.GORM, updatedCustomer)
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

	customer, err := ReadOneCustomer(db.GORM, customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Customer not found"})
		return
	}

	err = DeleteOneCustomer(db.GORM, *customer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": customer})
}
