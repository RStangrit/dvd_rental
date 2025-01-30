package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateStore(newStore *models.Store) error {
	return db.GORM.Table("store").Create(&newStore).Error
}

func ReadAllStores(pagination db.Pagination) ([]models.Store, int64, error) {
	var stores []models.Store
	var totalRecords int64

	db.GORM.Table("store").Count(&totalRecords)
	err := db.GORM.Table("store").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("store_id asc").Find(&stores).Error
	return stores, totalRecords, err
}

func ReadOneStore(storeID int64) (*models.Store, error) {
	var store models.Store
	err := db.GORM.Table("store").First(&store, storeID).Error
	return &store, err
}

func UpdateOneStore(store models.Store) error {
	return db.GORM.Table("store").Where("store_id = ?", store.StoreID).Updates(store).Error
}

func DeleteOneStore(store models.Store) error {
	return db.GORM.Table("store").Where("store_id = ?", store.StoreID).Delete(&store).Error
}
