package actor

import (
	"database/sql"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var mockDB *sql.DB
var gormDB *gorm.DB
var sqlMock sqlmock.Sqlmock

func TestMain(m *testing.M) {
	var err error

	mockDB, sqlMock, err = sqlmock.New()
	if err != nil {
		panic("Error creating sqlmock DB: " + err.Error())
	}

	gormDB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: mockDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("Failed to initialize GORM: " + err.Error())
	}

	code := m.Run()

	mockDB.Close()

	os.Exit(code)
}
