package user

import (
	"main/internal/models"
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID     int64          `json:"user_id" gorm:"type: integer; primaryKey;autoIncrement;not null"`
	Email      string         `json:"email" gorm:"type: varchar(45);not null"`
	Password   string         `json:"password" gorm:"type: varchar(60);not null"`
	LastUpdate time.Time      `json:"last_update" gorm:"type: timestamp;not null; autoUpdateTime;default:now()"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func init() {
	models.RegisterModel(&User{})
}
