package payment

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service *PaymentService
}

func NewPaymentHandler(service *PaymentService) *PaymentHandler {
	return &PaymentHandler{service: service}
}

func (handler *PaymentHandler) PostPaymentHandler(context *gin.Context) {
	var newPayment Payment
	if err := context.ShouldBindJSON(&newPayment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := handler.service.CreatePayment(&newPayment); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newPayment})
}

func (handler *PaymentHandler) GetPaymentsHandler(context *gin.Context) {
	var pagination db.Pagination
	if err := context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	payments, totalRecords, err := handler.service.ReadAllPayments(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": payments, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *PaymentHandler) GetPaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := handler.service.ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": payment})
}

func (handler *PaymentHandler) PutPaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := handler.service.ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedPayment Payment
	if err := context.ShouldBindJSON(&updatedPayment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data format"})
		return
	}

	updatedPayment.PaymentID = payment.PaymentID

	if err := handler.service.UpdateOnePayment(&updatedPayment); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedPayment})
}

func (handler *PaymentHandler) DeletePaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := handler.service.ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Payment not found"})
		return
	}

	if err := handler.service.DeleteOnePayment(payment); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": payment})
}
