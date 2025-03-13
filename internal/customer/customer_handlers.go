package customer

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerHandler struct {
	service *CustomerService
}

func NewCustomerHandler(service *CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

func (handler *CustomerHandler) PostCustomerHandler(context *gin.Context) {
	var newCustomer Customer
	var err error

	if err = context.ShouldBindJSON(&newCustomer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateCustomer(&newCustomer); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newCustomer})
}

func (handler *CustomerHandler) GetCustomersHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	customers, totalRecords, err := handler.service.ReadAllCustomers(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customers, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *CustomerHandler) GetCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := handler.service.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": customer})
}

func (handler *CustomerHandler) PutCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := handler.service.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedCustomer *Customer
	err = context.ShouldBindJSON(&updatedCustomer)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer data format"})
		return
	}

	updatedCustomer.CustomerID = customer.CustomerID

	err = handler.service.UpdateOneCustomer(updatedCustomer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedCustomer})
}

func (handler *CustomerHandler) DeleteCustomerHandler(context *gin.Context) {
	customerId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID format"})
		return
	}

	customer, err := handler.service.ReadOneCustomer(customerId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = handler.service.DeleteOneCustomer(customer)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete customer"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": customer})
}
