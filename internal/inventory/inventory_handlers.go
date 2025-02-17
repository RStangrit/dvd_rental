package inventory

import (
	"main/pkg/db"
	"main/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func PostInventoryHandler(context *gin.Context) {
	var newInventory Inventory
	var err error

	if err = context.ShouldBindJSON(&newInventory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err = ValidateInventory(&newInventory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newInventory.LastUpdate = time.Now()

	if err = CreateInventory(db.GORM, &newInventory); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": newInventory})
}

func GetInventoriesHandler(context *gin.Context) {
	var pagination db.Pagination
	var err error

	if err = context.ShouldBindQuery(&pagination); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pagination parameters"})
		return
	}

	inventories, totalRecords, err := ReadAllInventories(db.GORM, pagination)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inventories, "page": pagination.Page, "limit": pagination.Limit, "total": totalRecords})
}

func GetInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	inventory, err := ReadOneInventory(db.GORM, inventoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": inventory})
}

func PutInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	inventory, err := ReadOneInventory(db.GORM, inventoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updatedInventory Inventory
	err = context.ShouldBindJSON(&updatedInventory)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory data format"})
		return
	}

	if err = ValidateInventory(&updatedInventory); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedInventory.InventoryID = inventory.InventoryID
	updatedInventory.LastUpdate = time.Now()

	err = UpdateOneInventory(db.GORM, updatedInventory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": updatedInventory})
}

func DeleteInventoryHandler(context *gin.Context) {
	inventoryId, err := utils.GetIntParam(context, "id")
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid inventory ID format"})
		return
	}

	inventory, err := ReadOneInventory(db.GORM, inventoryId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = DeleteOneInventory(db.GORM, *inventory)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"deleted": inventory})
}
