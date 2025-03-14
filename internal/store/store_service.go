package store

import (
	"errors"
	"fmt"
	"main/pkg/db"
)

type StoreService struct {
	repo *StoreRepository
}

func NewStoreService(repo *StoreRepository) *StoreService {
	return &StoreService{repo: repo}
}

func (service *StoreService) CreateStore(newStore *Store) error {
	if err := service.ValidateStore(newStore); err != nil {
		return err
	}
	return service.repo.InsertStore(newStore)
}

func (service *StoreService) ReadAllStores(pagination db.Pagination) ([]Store, int64, error) {
	stores, totalRecords, err := service.repo.SelectAllStores(pagination)
	if err != nil {
		return nil, 0, err
	}
	return stores, totalRecords, nil
}

func (service *StoreService) ReadOneStore(storeID int64) (*Store, error) {
	store, err := service.repo.SelectOneStore(storeID)
	if err != nil {
		return nil, err
	}
	if store == nil {
		return nil, fmt.Errorf("store not found")
	}
	return store, nil
}

func (service *StoreService) UpdateOneStore(store *Store) error {
	if err := service.ValidateStore(store); err != nil {
		return err
	}
	return service.repo.UpdateOneStore(*store)
}

func (service *StoreService) DeleteOneStore(store *Store) error {
	return service.repo.DeleteOneStore(*store)
}

func (service *StoreService) ValidateStore(store *Store) error {
	if store.ManagerStaffID <= 0 {
		return errors.New("manager_staff_id must be a positive integer")
	}
	if store.AddressID <= 0 {
		return errors.New("address_id must be a positive integer")
	}
	return nil
}
