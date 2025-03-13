package user

import (
	"main/pkg/auth"
	"main/pkg/db"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, newUser *User) error {
	var err error
	newUser.Password, err = auth.GenerateHashFromPassword(newUser.Password)
	if err != nil {
		return err
	}

	return db.Table("user").Create(&newUser).Error
}

func ReadAllUsers(db *gorm.DB, pagination db.Pagination) ([]User, int64, error) {
	var users []User
	var totalRecords int64

	db.Table("user").Where("deleted_at IS NULL").Count(&totalRecords)
	err := db.Table("user").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("User_id asc").Find(&users).Error
	return users, totalRecords, err
}

func ReadOneUserById(db *gorm.DB, userID int64) (*User, error) {
	var user User
	err := db.Table("user").First(&user, userID).Error
	return &user, err
}

func ReadOneUserByEmail(db *gorm.DB, userEmail string) (User, error) {
	var user User
	err := db.Table("user").Where("email = ?", userEmail).First(&user).Error
	return user, err
}

func UpdateOneUser(db *gorm.DB, user User) error {
	return db.Table("user").Omit("User_id").Updates(user).Error
}

func DeleteOneUser(db *gorm.DB, user User) error {
	return db.Delete(&user).Error
}
