package store

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type StoreRepository struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) *StoreRepository {
	return &StoreRepository{db: db}
}

func (repo *StoreRepository) InsertStore(newStore *Store) error {
	return repo.db.Table("store").Create(&newStore).Error
}

func (repo *StoreRepository) SelectAllStores(pagination db.Pagination) ([]Store, int64, error) {
	var stores []Store
	var totalRecords int64

	repo.db.Table("store").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("store").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("store_id asc").
		Find(&stores).Error

	return stores, totalRecords, err
}

func (repo *StoreRepository) SelectOneStore(storeID int64) (*Store, error) {
	var store Store
	err := repo.db.Table("store").First(&store, storeID).Error
	return &store, err
}

func (repo *StoreRepository) UpdateOneStore(store Store) error {
	return repo.db.Table("store").Where("store_id = ?", store.StoreID).Updates(store).Error
}

func (repo *StoreRepository) DeleteOneStore(store Store) error {
	return repo.db.Table("store").Where("store_id = ?", store.StoreID).Delete(&store).Error
}
