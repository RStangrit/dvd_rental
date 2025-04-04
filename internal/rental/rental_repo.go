package rental

import (
	"main/pkg/db"

	"gorm.io/gorm"
)

// Dependency Injection Principle violated here, needs to be rewritten to interfaces
type RentalRepository struct {
	db *gorm.DB
}

func NewRentalRepository(db *gorm.DB) *RentalRepository {
	return &RentalRepository{db: db}
}

func (repo *RentalRepository) InsertRental(newRental *Rental) error {
	return repo.db.Table("rental").Create(&newRental).Error
}

func (repo *RentalRepository) SelectAllRentals(pagination db.Pagination) ([]Rental, int64, error) {
	var rentals []Rental
	var totalRecords int64

	repo.db.Table("rental").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("rental").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("rental_id asc").
		Find(&rentals).Error

	return rentals, totalRecords, err
}

func (repo *RentalRepository) SelectOneRental(rentalID int64) (*Rental, error) {
	var rental Rental
	err := repo.db.Table("rental").First(&rental, rentalID).Error
	return &rental, err
}

func (repo *RentalRepository) UpdateOneRental(rental Rental) error {
	return repo.db.Table("rental").Where("rental_id = ?", rental.RentalID).Updates(rental).Error
}

func (repo *RentalRepository) DeleteOneRental(rental Rental) error {
	return repo.db.Table("rental").Where("rental_id = ?", rental.RentalID).Delete(&rental).Error
}
