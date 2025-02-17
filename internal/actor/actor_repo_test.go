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

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery(`INSERT INTO "actor" (.+) RETURNING`).
		WithArgs(newActor.FirstName, newActor.LastName, sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id"}).AddRow(1))
	sqlMock.ExpectCommit()

	err := CreateActor(gormDB, newActor)

	assert.NoError(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_ReadAllActors(t *testing.T) {
	expectedActors := []*Actor{
		{ActorID: 1, FirstName: "John", LastName: "Doe"},
		{ActorID: 2, FirstName: "Jane", LastName: "Smith"},
	}

	sqlMock.ExpectQuery(`SELECT count\(\*\) FROM "actor"`).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	sqlMock.ExpectQuery(`SELECT \* FROM "actor" WHERE "actor"."deleted_at" IS NULL ORDER BY actor_id asc LIMIT \$1`).
		WithArgs(10). // Pagination: limit=10
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name"}).
			AddRow(expectedActors[0].ActorID, expectedActors[0].FirstName, expectedActors[0].LastName).
			AddRow(expectedActors[1].ActorID, expectedActors[1].FirstName, expectedActors[1].LastName))

	pagination := db.Pagination{Page: 1, Limit: 10}

	actors, total, err := ReadAllActors(gormDB, pagination)

	assert.NoError(t, err)
	assert.Equal(t, int64(2), total)
	assert.Equal(t, expectedActors, actors)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
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
	sqlMock.ExpectQuery(`SELECT \* FROM "actor" WHERE "actor"."actor_id" = \$1 AND "actor"."deleted_at" IS NULL ORDER BY "actor"."actor_id" LIMIT \$2`).
		WithArgs(expectedActor.ActorID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name", "last_update", "deleted_at"}).
			AddRow(expectedActor.ActorID, expectedActor.FirstName, expectedActor.LastName, expectedActor.LastUpdate, expectedActor.DeletedAt))
	actor, err := ReadOneActor(gormDB, int64(expectedActor.ActorID))
	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)
	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_ReadOneActorFilms(t *testing.T) {
	fixedTime := time.Now()

	expectedFilmActor := &film_actor.FilmActor{
		ActorID:    1,
		FilmID:     1,
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	expectedActor := &Actor{
		ActorID:   1,
		FirstName: "John",
		LastName:  "Doe",
		ActorFilms: []film_actor.FilmActor{
			*expectedFilmActor,
		},
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	sqlMock.ExpectQuery(`SELECT \* FROM "actor" WHERE actor.actor_id = \$1 AND "actor"."deleted_at" IS NULL ORDER BY "actor"."actor_id" LIMIT \$2`).
		WithArgs(expectedActor.ActorID, 1).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "first_name", "last_name", "film_actor.actor_id", "film_actor.film_id", "film_actor.last_update", "film_actor.deleted_at", "last_update", "deleted_at"}).
			AddRow(expectedActor.ActorID, expectedActor.FirstName, expectedActor.LastName, expectedFilmActor.ActorID, expectedFilmActor.FilmID, expectedFilmActor.LastUpdate, expectedFilmActor.DeletedAt, expectedActor.LastUpdate, expectedActor.DeletedAt))

	sqlMock.ExpectQuery(`SELECT \* FROM "film_actor" WHERE "film_actor"."actor_id" = \$1 AND "film_actor"."deleted_at" IS NULL`).
		WillReturnRows(sqlmock.NewRows([]string{"actor_id", "film_id", "last_update", "deleted_at"}).AddRow(expectedFilmActor.ActorID, expectedFilmActor.FilmID, expectedFilmActor.LastUpdate, expectedFilmActor.DeletedAt))

	actor, err := ReadOneActorFilms(gormDB, int64(expectedActor.ActorID))

	assert.NoError(t, err)
	assert.Equal(t, expectedActor, actor)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}

func Test_UpdateOneActor(t *testing.T) {
	fixedTime := time.Now()

	actor := &Actor{
		ActorID:    1,
		FirstName:  "John",
		LastName:   "Doe",
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "actor" SET .+`).
		WithArgs(actor.FirstName, actor.LastName, sqlmock.AnyArg(), actor.ActorID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	sqlMock.ExpectCommit()

	err := UpdateOneActor(gormDB, *actor)

	assert.NoError(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}

}

func Test_DeleteOneActor(t *testing.T) {
	// return db.Delete(&actor).Error
	fixedTime := time.Now()

	actor := &Actor{
		ActorID:    1,
		FirstName:  "John",
		LastName:   "Doe",
		LastUpdate: fixedTime,
		DeletedAt:  gorm.DeletedAt{Valid: false},
	}

	sqlMock.ExpectBegin()
	sqlMock.ExpectExec(`UPDATE "actor" SET .+`).
		WithArgs(sqlmock.AnyArg(), actor.ActorID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	sqlMock.ExpectCommit()

	err := DeleteOneActor(gormDB, *actor)

	assert.NoError(t, err)

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %v", err)
	}
}
