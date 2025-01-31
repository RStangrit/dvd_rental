package view

type SalesByFilmCategory struct {
	Category   string  `json:"category" gorm:"type:varchar(25);not null"`
	TotalSales float64 `json:"total_sales" gorm:"type:numeric;not null"`
}

func (SalesByFilmCategory) TableName() string {
	return "sales_by_film_category"
}
