package staff

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateStaff(db *gorm.DB, newStaff *Staff) error {
	return db.Table("staff").Create(&newStaff).Error
}

func ReadAllStaff(db *gorm.DB, pagination db.Pagination) ([]Staff, int64, error) {
	var staffList []Staff
	var totalRecords int64

	db.Table("staff").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("staff").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("staff_id asc").Find(&staffList).Error
	return staffList, totalRecords, err
}

func ReadOneStaff(db *gorm.DB, staffId int64) (*Staff, error) {
	var staff Staff
	err := db.Table("staff").First(&staff, staffId).Error
	return &staff, err
}

func UpdateOneStaff(db *gorm.DB, staff Staff) error {
	return db.Table("staff").Omit("staff_id").Updates(staff).Error
}

func DeleteOneStaff(db *gorm.DB, staff Staff) error {
	return db.Delete(&staff).Error
}
