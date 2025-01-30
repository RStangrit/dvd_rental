package services

import (
	"errors"
	"main/internal/models"
)

func ValidateInventory(inventory *models.Inventory) error {
	if inventory.FilmID <= 0 {
		return errors.New("film_id must be a positive integer")
	}
	if inventory.StoreID <= 0 {
		return errors.New("store_id must be a positive integer")
	}
	return nil
}
