package actor

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Test_CreateActor(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Fatalf("Error connecting to GORM: %v", err)
	}

	newActor := &Actor{
		FirstName: "John",
		LastName:  "Doe",
		DeletedAt: gorm.DeletedAt{Valid: false},
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "actor" (.+) RETURNING`).
		WithArgs(newActor.FirstName, newActor.LastName, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id"}).AddRow(1))
	mock.ExpectCommit()

	err = CreateActor(gormDB, newActor)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("Something went wrong: %v", err)
	}
}
