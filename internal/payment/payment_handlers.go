package payment

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPaymentHandler(context *gin.Context) {
	var newPayment Payment
	var err error

	if err = context.ShouldBindJSON(&newPayment); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreatePayment(&newPayment); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newPayment})
}

func GetPaymentsHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	payments, totalRecords, err := ReadAllPayments(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": payments, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetPaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": payment})
}

func PutPaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedPayment Payment
	err = context.ShouldBindJSON(&updatedPayment)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment data format"})
		return
	}

	updatedPayment.PaymentID = payment.PaymentID

	err = UpdateOnePayment(updatedPayment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedPayment})
}

func DeletePaymentHandler(context *gin.Context) {
	paymentID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payment ID format"})
		return
	}

	payment, err := ReadOnePayment(paymentID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Payment not found"})
		return
	}

	err = DeleteOnePayment(*payment)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete payment"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": payment})
}
