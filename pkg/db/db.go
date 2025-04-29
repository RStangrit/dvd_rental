package db

import (
	"fmt"
	"log"
	"main/config"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	GORM *gorm.DB
)

func InitDb() *gorm.DB {
	params := config.LoadConfig()
	dsn := params.DSN
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  true,          // Disable color
		},
	)
	maxRetries := 5
	retryInterval := 3 * time.Second

	connectWithRetry(dsn, newLogger, maxRetries, retryInterval)

	pool, err := GORM.DB()
	if err != nil {
		panic(err)
	}

	pool.SetMaxOpenConns(10)
	pool.SetMaxIdleConns(2)
	pool.SetConnMaxLifetime(30 * time.Minute)
	pool.SetConnMaxIdleTime(10 * time.Minute)
	fmt.Println("database has been successfully configured")

	trackQueryTime()

	return GORM
}

func trackQueryTime() {
	GORM.Callback().Query().Before("gorm:query").Register("start_time", func(db *gorm.DB) {
		db.InstanceSet("start_time", time.Now())
	})

	GORM.Callback().Query().After("gorm:query").Register("end_time", func(db *gorm.DB) {
		startTime, ok := db.InstanceGet("start_time")
		if !ok {
			return
		}

		duration := time.Since(startTime.(time.Time))
		fmt.Printf("Query took: %v\n", duration)
	})
}

func connectWithRetry(dsn string, newLogger logger.Interface, maxRetries int, retryInterval time.Duration) {
	var err error
	for i := 0; i < maxRetries; i++ {
		GORM, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err == nil {
			fmt.Println("Connection to the database has been successfully established")
			return
		}
		fmt.Printf("Failed to connect to the database (attempt %d/%d): %v\n", i+1, maxRetries, err)
		time.Sleep(retryInterval)
	}
	panic(fmt.Sprintf("Could not connect to the database after %d attempts: %v", maxRetries, err))
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
