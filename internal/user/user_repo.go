package user

import (
	"main/pkg/auth"
	"main/pkg/db"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) InsertUser(newUser *User) error {
	var err error
	newUser.Password, err = auth.GenerateHashFromPassword(newUser.Password)
	if err != nil {
		return err
	}
	return repo.db.Table("user").Create(&newUser).Error
}

func (repo *UserRepository) SelectAllUsers(pagination db.Pagination) ([]User, int64, error) {
	var users []User
	var totalRecords int64

	repo.db.Table("user").Where("deleted_at IS NULL").Count(&totalRecords)
	err := repo.db.Table("user").
		Offset(pagination.GetOffset()).
		Limit(pagination.GetLimit()).
		Order("user_id asc").
		Find(&users).Error

	return users, totalRecords, err
}

func (repo *UserRepository) SelectOneUserById(userID int64) (*User, error) {
	var user User
	err := repo.db.Table("user").First(&user, userID).Error
	return &user, err
}

func (repo *UserRepository) SelectOneUserByEmail(userEmail string) (*User, error) {
	var user *User
	err := repo.db.Table("user").Where("email = ?", userEmail).First(&user).Error
	return user, err
}

func (repo *UserRepository) UpdateOneUser(user User) error {
	return repo.db.Table("user").Omit("user_id").Updates(user).Error
}

func (repo *UserRepository) DeleteOneUser(user User) error {
	return repo.db.Delete(&user).Error
}
