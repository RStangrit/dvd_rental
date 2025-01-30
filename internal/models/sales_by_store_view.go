package models

type SalesByStore struct {
	Store      string  `json:"store" gorm:"type:text;not null"`
	Manager    string  `json:"manager" gorm:"type:text;not null"`
	TotalSales float64 `json:"total_sales" gorm:"type:numeric;not null"`
}

func (SalesByStore) TableName() string {
	return "sales_by_store"
}
