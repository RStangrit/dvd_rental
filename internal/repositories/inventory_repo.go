package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateInventory(inventory *models.Inventory) error {
	return db.GORM.Table("inventory").Create(&inventory).Error
}

func ReadAllInventories(pagination db.Pagination) ([]models.Inventory, int64, error) {
	var inventories []models.Inventory
	var totalRecords int64

	query := db.GORM.Model(&models.Inventory{})
	query.Count(&totalRecords)

	offset := (pagination.Page - 1) * pagination.Limit
	err := query.Limit(pagination.Limit).Offset(offset).Find(&inventories).Error
	return inventories, totalRecords, err
}

func ReadOneInventory(inventoryID int64) (*models.Inventory, error) {
	var inventory models.Inventory
	err := db.GORM.First(&inventory, inventoryID).Error
	return &inventory, err
}

func UpdateOneInventory(inventory models.Inventory) error {
	return db.GORM.Save(&inventory).Error
}

func DeleteOneInventory(inventory models.Inventory) error {
	return db.GORM.Delete(&inventory).Error
}
