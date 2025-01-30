package handlers

import (
	"main/internal/models"
	"main/internal/repositories"
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostStoreHandler(context *gin.Context) {
	var newStore models.Store
	var err error

	if err = context.ShouldBindJSON(&newStore); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = repositories.CreateStore(&newStore); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newStore})
}

func GetStoresHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	stores, totalRecords, err := repositories.ReadAllStores(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": stores, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := repositories.ReadOneStore(storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": store})
}

func PutStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := repositories.ReadOneStore(storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedStore models.Store
	err = context.ShouldBindJSON(&updatedStore)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store data format"})
		return
	}

	updatedStore.StoreID = store.StoreID

	err = repositories.UpdateOneStore(updatedStore)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedStore})
}

func DeleteStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := repositories.ReadOneStore(storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Store not found"})
		return
	}

	err = repositories.DeleteOneStore(*store)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": store})
}
