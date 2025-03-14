package staff

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

type StaffRepository struct {
	db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
	return &StaffRepository{db: db}
}

func (repo *StaffRepository) InsertStaff(newStaff *Staff) error {
	return repo.db.Table("staff").Create(&newStaff).Error
}

func (repo *StaffRepository) SelectAllStaffs(pagination db.Pagination) ([]Staff, int64, error) {
	var staffList []Staff
	var totalRecords int64

	repo.db.Table("staff").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("staff").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("staff_id asc").
		Find(&staffList).Error

	return staffList, totalRecords, err
}

func (repo *StaffRepository) SelectOneStaff(staffId int64) (*Staff, error) {
	var staff Staff
	err := repo.db.Table("staff").First(&staff, staffId).Error
	return &staff, err
}

func (repo *StaffRepository) UpdateOneStaff(staff Staff) error {
	return repo.db.Table("staff").Omit("staff_id").Updates(staff).Error
}

func (repo *StaffRepository) DeleteOneStaff(staff Staff) error {
	return repo.db.Delete(&staff).Error
}
