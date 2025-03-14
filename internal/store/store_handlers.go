package store

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type StoreHandler struct {
	service *StoreService
}

func NewStoreHandler(service *StoreService) *StoreHandler {
	return &StoreHandler{service: service}
}

func (handler *StoreHandler) PostStoreHandler(context *gin.Context) {
	var newStore Store
	var err error

	if err = context.ShouldBindJSON(&newStore); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateStore(&newStore); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newStore})
}

func (handler *StoreHandler) GetStoresHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	stores, totalRecords, err := handler.service.ReadAllStores(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": stores, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *StoreHandler) GetStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := handler.service.ReadOneStore(storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": store})
}

func (handler *StoreHandler) PutStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := handler.service.ReadOneStore(storeID)
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

	err = handler.service.UpdateOneStore(&updatedStore)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update store"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedStore})
}

func (handler *StoreHandler) DeleteStoreHandler(context *gin.Context) {
	storeID, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid store ID format"})
		return
	}

	store, err := handler.service.ReadOneStore(storeID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Store not found"})
		return
	}

	err = handler.service.DeleteOneStore(store)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete store"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": store})
}
