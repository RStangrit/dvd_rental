package actor

import (
	"main/internal/film_actor"
	"main/pkg/db"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_CreateActor(t *testing.T) {
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

	err := CreateActor(gormDB, newActor)

	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_ReadAllActors(t *testing.T) {
	expectedActors := []Actor{
		{ActorID: 1, FirstName: "John", LastName: "Doe"},
		{ActorID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	mock.ExpectQuery(`SELECT count\(\*\) FROM "actor"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery(`SELECT \* FROM "actor" WHERE "actor"."deleted_at" IS NULL ORDER BY actor_id asc LIMIT \$1`).
		WithArgs(10). // Pagination: limit=10
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name"}).
			AddRow(expectedActors[0].ActorID, expectedActors[0].FirstName, expectedActors[0].LastName).
			AddRow(expectedActors[1].ActorID, expectedActors[1].FirstName, expectedActors[1].LastName))

	pagination := db.Pagination{Page: 1, Limit: 10}

	actors, total, err := ReadAllActors(gormDB, pagination)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Equal(t, expectedActors, actors)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_ReadOneActor(t *testing.T) {
	expectedActor := &Actor{
		ActorID:    1,
		FirstName:  "John",
		LastName:   "Doe",
		LastUpdate: time.Now(),
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}
	mock.ExpectQuery(`SELECT \* FROM "actor" WHERE "actor"."actor_id" = \$1 AND "actor"."deleted_at" IS NULL ORDER BY "actor"."actor_id" LIMIT \$2`).
		WithArgs(expectedActor.ActorID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name", "last_update", "deleted_at"}).
			AddRow(expectedActor.ActorID, expectedActor.FirstName, expectedActor.LastName, expectedActor.LastUpdate, expectedActor.DeletedAt))
	actor, err := ReadOneActor(gormDB, int64(expectedActor.ActorID))
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_ReadOneActorFilms(t *testing.T) {
	fixedTime := time.Date(2025, time.February, 14, 9, 56, 28, 693128929, time.UTC)

	expectedFilmActor := &film_actor.FilmActor{
		ActorID:    1,
		FilmID:     1,
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	expectedActor := Actor{
		ActorID:   1,
		FirstName: "John",
		LastName:  "Doe",
		ActorFilms: []film_actor.FilmActor{
			*expectedFilmActor,
		},
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	mock.ExpectQuery(`SELECT \* FROM "actor" WHERE actor.actor_id = \$1 AND "actor"."deleted_at" IS NULL ORDER BY "actor"."actor_id" LIMIT \$2`).
		WithArgs(expectedActor.ActorID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name", "film_actor.actor_id", "film_actor.film_id", "film_actor.last_update", "film_actor.deleted_at", "last_update", "deleted_at"}).
			AddRow(expectedActor.ActorID, expectedActor.FirstName, expectedActor.LastName, expectedFilmActor.ActorID, expectedFilmActor.FilmID, expectedFilmActor.LastUpdate, expectedFilmActor.DeletedAt, expectedActor.LastUpdate, expectedActor.DeletedAt))

	mock.ExpectQuery(`SELECT \* FROM "film_actor" WHERE "film_actor"."actor_id" = \$1 AND "film_actor"."deleted_at" IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "film_id", "last_update", "deleted_at"}).AddRow(expectedFilmActor.ActorID, expectedFilmActor.FilmID, expectedFilmActor.LastUpdate, expectedFilmActor.DeletedAt))

	actor, err := ReadOneActorFilms(gormDB, int64(expectedActor.ActorID))

	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
