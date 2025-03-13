package inventory

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type InventoryRepository struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) *InventoryRepository {
	return &InventoryRepository{db: db}
}

func (repo *InventoryRepository) InsertInventory(newInventory *Inventory) error {
	return repo.db.Table("inventory").Create(&newInventory).Error
}

func (repo *InventoryRepository) SelectAllInventories(pagination db.Pagination) ([]Inventory, int64, error) {
	var inventories []Inventory
	var totalRecords int64

	repo.db.Table("inventory").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("inventory").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("inventory_id asc").Find(&inventories).Error
	return inventories, totalRecords, err
}

func (repo *InventoryRepository) SelectOneInventory(inventoryID int64) (*Inventory, error) {
	var inventory Inventory
	err := repo.db.Table("inventory").First(&inventory, inventoryID).Error
	return &inventory, err
}

func (repo *InventoryRepository) UpdateOneInventory(inventory Inventory) error {
	return repo.db.Table("inventory").Omit("inventory_id").Updates(inventory).Error
}

func (repo *InventoryRepository) DeleteOneInventory(inventory Inventory) error {
	return repo.db.Delete(&inventory).Error
}
