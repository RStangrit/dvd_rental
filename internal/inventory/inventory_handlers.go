package inventory

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type InventoryHandler struct {
	service *InventoryService
}

func NewInventoryHandler(service *InventoryService) *InventoryHandler {
	return &InventoryHandler{service: service}
}

func (handler *InventoryHandler) PostInventoryHandler(context *gin.Context) {
	var newInventory Inventory
	var err error

	if err = context.ShouldBindJSON(&newInventory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = handler.service.CreateInventory(&newInventory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newInventory})
}

func (handler *InventoryHandler) GetInventoriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	inventories, totalRecords, err := handler.service.ReadAllInventories(pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inventories, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func (handler *InventoryHandler) GetInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	inventory, err := handler.service.ReadOneInventory(inventoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inventory})
}

func (handler *InventoryHandler) PutInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	var updatedInventory Inventory
	if err = context.ShouldBindJSON(&updatedInventory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory data format"})
		return
	}

	updatedInventory.InventoryID = int(inventoryId)
	updatedInventory.LastUpdate = time.Now()

	if err = handler.service.UpdateOneInventory(&updatedInventory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedInventory})
}

func (handler *InventoryHandler) DeleteInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	var deletedInventory Inventory
	deletedInventory.InventoryID = int(inventoryId)

	if err = handler.service.DeleteOneInventory(&deletedInventory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": deletedInventory})
}
