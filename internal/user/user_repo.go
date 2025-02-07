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

func ReadAllUsers(pagination db.Pagination) ([]User, int64, error) {
	var Users []User
	var totalRecords int64

	db.GORM.Table("User").Count(&totalRecords)
	err := db.GORM.Table("user").Offset(pagination.GetOffset()).Limit(pagination.GetLimit()).Order("User_id asc").Find(&Users).Error
	return Users, totalRecords, err
}

func ReadOneUser(UserId int64) (*User, error) {
	var User User
	err := db.GORM.Table("user").First(&User, UserId).Error
	return &User, err
}

func UpdateOneUser(User User) error {
	return db.GORM.Table("user").Omit("User_id").Updates(User).Error
}

func DeleteOneUser(User User) error {
	return db.GORM.Delete(&User).Error
}
