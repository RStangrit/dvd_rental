package inventory

import (
	"main/pkg/db"
)

func CreateInventory(inventory *Inventory) error {
	return db.GORM.Table("inventory").Create(&inventory).Error
}

func ReadAllInventories(pagination db.Pagination) ([]Inventory, int64, error) {
	var inventories []Inventory
	var totalRecords int64

	query := db.GORM.Model(&Inventory{})
	query.Count(&totalRecords)

	offset := (pagination.Page - 1) * pagination.Limit
	err := query.Limit(pagination.Limit).Offset(offset).Find(&inventories).Error
	return inventories, totalRecords, err
}

func ReadOneInventory(inventoryID int64) (*Inventory, error) {
	var inventory Inventory
	err := db.GORM.First(&inventory, inventoryID).Error
	return &inventory, err
}

func UpdateOneInventory(inventory Inventory) error {
	return db.GORM.Save(&inventory).Error
}

func DeleteOneInventory(inventory Inventory) error {
	return db.GORM.Delete(&inventory).Error
}
