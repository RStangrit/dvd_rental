package inventory

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateInventory(db *gorm.DB, inventory *Inventory) error {
	return db.Table("inventory").Create(&inventory).Error
}

func ReadAllInventories(db *gorm.DB, pagination db.Pagination) ([]Inventory, int64, error) {
	var inventories []Inventory
	var totalRecords int64

	query := db.Model(&Inventory{})
	query.Where("deleted_at IS NULL").Count(&totalRecords)

	offset := (pagination.Page - 1) * pagination.Limit
	err := query.Limit(pagination.Limit).Offset(offset).Find(&inventories).Error
	return inventories, totalRecords, err
}

func ReadOneInventory(db *gorm.DB, inventoryID int64) (*Inventory, error) {
	var inventory Inventory
	err := db.First(&inventory, inventoryID).Error
	return &inventory, err
}

func UpdateOneInventory(db *gorm.DB, inventory Inventory) error {
	return db.Save(&inventory).Error
}

func DeleteOneInventory(db *gorm.DB, inventory Inventory) error {
	return db.Delete(&inventory).Error
}
