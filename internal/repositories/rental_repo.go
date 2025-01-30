package repositories

import (
	"main/internal/models"
	"main/pkg/db"
)

func CreateRental(newRental *models.Rental) error {
	return db.GORM.Table("rental").Create(&newRental).Error
}

func ReadAllRentals(pagination db.Pagination) ([]models.Rental, int64, error) {
	var rentals []models.Rental
	var totalRecords int64

	db.GORM.Table("rental").Count(&totalRecords)
	err := db.GORM.Table("rental").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("rental_id asc").Find(&rentals).Error
	return rentals, totalRecords, err
}

func ReadOneRental(rentalID int64) (*models.Rental, error) {
	var rental models.Rental
	err := db.GORM.Table("rental").First(&rental, rentalID).Error
	return &rental, err
}

func UpdateOneRental(rental models.Rental) error {
	return db.GORM.Table("rental").Where("rental_id = ?", rental.RentalID).Updates(rental).Error
}

func DeleteOneRental(rental models.Rental) error {
	return db.GORM.Table("rental").Where("rental_id = ?", rental.RentalID).Delete(&rental).Error
}
