package rental

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateRental(db *gorm.DB, newRental *Rental) error {
	return db.Table("rental").Create(&newRental).Error
}

func ReadAllRentals(db *gorm.DB, pagination db.Pagination) ([]Rental, int64, error) {
	var rentals []Rental
	var totalRecords int64

	db.Table("rental").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("rental").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("rental_id asc").Find(&rentals).Error
	return rentals, totalRecords, err
}

func ReadOneRental(db *gorm.DB, rentalID int64) (*Rental, error) {
	var rental Rental
	err := db.Table("rental").First(&rental, rentalID).Error
	return &rental, err
}

func UpdateOneRental(db *gorm.DB, rental Rental) error {
	return db.Table("rental").Where("rental_id = ?", rental.RentalID).Updates(rental).Error
}

func DeleteOneRental(db *gorm.DB, rental Rental) error {
	return db.Table("rental").Where("rental_id = ?", rental.RentalID).Delete(&rental).Error
}
