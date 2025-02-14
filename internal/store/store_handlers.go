package store

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostStoreHandler(context *gin.Context) {
	var newStore Store
	var err error

	if err = context.ShouldBindJSON(&newStore); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = CreateStore(db.GORM, &newStore); err != nil {
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

	stores, totalRecords, err := ReadAllStores(db.GORM, pagination)
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

	store, err := ReadOneStore(db.GORM, storeID)
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

	store, err := ReadOneStore(db.GORM, storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedStore Store
	err = context.ShouldBindJSON(&updatedStore)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store data format"})
		return
	}

	updatedStore.StoreID = store.StoreID

	err = UpdateOneStore(db.GORM, updatedStore)
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

	store, err := ReadOneStore(db.GORM, storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Store not found"})
		return
	}

	err = DeleteOneStore(db.GORM, *store)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": store})
}
