package store

import (
	"errors"
)

func ValidateStore(store *Store) error {
	if store.StoreID <= 0 {
		return errors.New("store_id must be a positive integer")
	}

	if store.ManagerStaffID <= 0 {
		return errors.New("manager_staff_id must be a positive integer")
	}

	if store.AddressID <= 0 {
		return errors.New("address_id must be a positive integer")
	}

	return nil
}
