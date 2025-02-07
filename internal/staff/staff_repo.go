package staff

import (
	"main/pkg/db"
)

func CreateStaff(newStaff *Staff) error {
	return db.GORM.Table("staff").Create(&newStaff).Error
}

func ReadAllStaff(pagination db.Pagination) ([]Staff, int64, error) {
	var staffList []Staff
	var totalRecords int64

	db.GORM.Table("staff").Count(&totalRecords)
	err := db.GORM.Table("staff").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("staff_id asc").Find(&staffList).Error
	return staffList, totalRecords, err
}

func ReadOneStaff(staffId int64) (*Staff, error) {
	var staff Staff
	err := db.GORM.Table("staff").First(&staff, staffId).Error
	return &staff, err
}

func UpdateOneStaff(staff Staff) error {
	return db.GORM.Table("staff").Omit("staff_id").Updates(staff).Error
}

func DeleteOneStaff(staff Staff) error {
	return db.GORM.Delete(&staff).Error
}
