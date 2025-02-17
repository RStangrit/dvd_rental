package store

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateStore(db *gorm.DB, newStore *Store) error {
	return db.Table("store").Create(&newStore).Error
}

func ReadAllStores(db *gorm.DB, pagination db.Pagination) ([]Store, int64, error) {
	var stores []Store
	var totalRecords int64

	db.Table("store").Count(&totalRecords)
	err := db.Table("store").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("store_id asc").Find(&stores).Error
	return stores, totalRecords, err
}

func ReadOneStore(db *gorm.DB, storeID int64) (*Store, error) {
	var store Store
	err := db.Table("store").First(&store, storeID).Error
	return &store, err
}

func UpdateOneStore(db *gorm.DB, store Store) error {
	return db.Table("store").Where("store_id = ?", store.StoreID).Updates(store).Error
}

func DeleteOneStore(db *gorm.DB, store Store) error {
	return db.Table("store").Where("store_id = ?", store.StoreID).Delete(&store).Error
}
