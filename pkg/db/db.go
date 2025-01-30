package db

import (
	"fmt"
	"main/config"
	"main/internal/models"
	"main/internal/seeds"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	GORM *gorm.DB
)

func InitDb() error {
	params := config.LoadConfig()
	dsn := params.DSN

	var err error

	GORM, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("connection to the database has been successfully established")

	err = createTables()
	if err != nil {
		panic(err)
	}

	err = seedData(seeds.ReturnLangSeeds())
	if err != nil {
		panic(err)
	}

	return nil
}

func createTables() error {
	err := GORM.AutoMigrate(
		&models.Language{},
		&models.Actor{},
		&models.Film{},
		&models.Category{},
		&models.FilmActor{},
		&models.Inventory{},
		&models.FilmCategory{},
		&models.Country{},
	)

	return err
}

func seedData(languages []models.Language) error {
	if err := GORM.Create(&languages).Error; err != nil {
		return err
	}

	return nil
}

type Pagination struct {
	Page  int `form:"page" json:"page"`
	Limit int `form:"limit" json:"limit"`
}

func (p *Pagination) GetOffset() int {
	if p.Page < 1 {
		p.Page = 1
	}
	return (p.Page - 1) * p.Limit
}

func (p *Pagination) GetLimit() int {
	if p.Limit < 1 {
		p.Limit = 10
	}
	return p.Limit
}
