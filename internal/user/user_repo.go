package user

import (
	"fmt"
	"main/pkg/db"
	"main/pkg/utils"
)

func CreateUser(newUser *User) error {
	var err error
	fmt.Println(newUser)
	newUser.Password, err = utils.GenerateHashFromPassword(newUser.Password)
	if err != nil {
		return err
	}

	fmt.Println(newUser)
	return db.GORM.Table("user").Create(&newUser).Error
}

func ReadAllusers(pagination db.Pagination) ([]User, int64, error) {
	var users []User
	var totalRecords int64

	db.GORM.Table("user").Count(&totalRecords)
	err := db.GORM.Table("user").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("User_id asc").Find(&users).Error
	return users, totalRecords, err
}

func ReadOneUserById(userID int64) (*User, error) {
	var user User
	err := db.GORM.Table("user").First(&user, userID).Error
	return &user, err
}

func ReadOneUserByEmail(userEmail string) (User, error) {
	var user User
	err := db.GORM.Table("user").Where("email = ?", userEmail).First(&user).Error
	return user, err
}

func UpdateOneUser(user User) error {
	return db.GORM.Table("user").Omit("User_id").Updates(user).Error
}

func DeleteOneUser(user User) error {
	return db.GORM.Delete(&user).Error
}
