package inventory

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type InventoryService struct {
	repo *InventoryRepository
}

func NewInventoryService(repo *InventoryRepository) *InventoryService {
	return &InventoryService{repo: repo}
}

func (service *InventoryService) CreateInventory(newInventory *Inventory) error {
	err := service.ValidateInventory(newInventory)
	if err != nil {
		return err
	} else {
		return service.repo.InsertInventory(newInventory)
	}
}

func (service *InventoryService) CreateInventories(newInventories []*Inventory) ([]string, []Inventory, error) {
	var validationErrors []string
	var createdInventories []Inventory

	for _, newInventory := range newInventories {
		if err := service.ValidateInventory(newInventory); err != nil {
			validationErrors = append(validationErrors, err.Error())
			continue
		}

		if err := service.repo.InsertInventory(newInventory); err != nil {
			return validationErrors, createdInventories, err
		}

		createdInventories = append(createdInventories, *newInventory)
	}
	return validationErrors, createdInventories, nil
}

func (service *InventoryService) ReadAllInventories(pagination db.Pagination) ([]Inventory, int64, error) {
	inventories, totalRecords, err := service.repo.SelectAllInventories(pagination)
	if err != nil {
		return nil, 0, err
	}
	return inventories, totalRecords, nil
}

func (service *InventoryService) ReadOneInventory(inventoryID int64) (*Inventory, error) {
	inventory, err := service.repo.SelectOneInventory(inventoryID)
	if err != nil {
		return nil, err
	}
	if inventory == nil {
		return nil, fmt.Errorf("Inventory not found")
	}
	return inventory, nil
}

func (service *InventoryService) UpdateOneInventory(inventory *Inventory) error {
	err := service.ValidateInventory(inventory)
	if err != nil {
		return err
	} else {
		return service.repo.UpdateOneInventory(*inventory)
	}
}

func (service *InventoryService) DeleteOneInventory(inventory *Inventory) error {
	return service.repo.DeleteOneInventory(*inventory)
}

var ErrInvalidFilmID = errors.New("film_id must be a positive integer")
var ErrInvalidStoreID = errors.New("store_id must be a positive integer")

func (service *InventoryService) ValidateInventory(inventory *Inventory) error {
	if inventory.FilmID <= 0 {
		return ErrInvalidFilmID
	}
	if inventory.StoreID <= 0 {
		return ErrInvalidStoreID
	}
	return nil
}
