package models

type Language struct {
	LanguageID int    `json:"language_id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	LastUpdate string `json:"last_update" gorm:"autoUpdateTime"`
}

func (Language) TableName() string {
	return "language"
}
