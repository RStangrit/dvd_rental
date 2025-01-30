package models

type CustomerList struct {
	ID      int    `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Name    string `json:"name" gorm:"type:text;not null"`
	Address string `json:"address" gorm:"type:varchar(50);not null"`
	ZipCode string `json:"zip_code" gorm:"type:varchar(10);not null"`
	Phone   string `json:"phone" gorm:"type:varchar(20);not null"`
	City    string `json:"city" gorm:"type:varchar(50);not null"`
	Country string `json:"country" gorm:"type:varchar(50);not null"`
	Notes   string `json:"notes" gorm:"type:text;not null"`
	SID     int16  `json:"sid" gorm:"type:int2;not null"`
}

func (CustomerList) TableName() string {
	return "customer_list"
}
