package migrations

import (
	"fmt"
	"main/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func launchMigrationsMigrate() error {
	params := config.LoadConfig()
	migrationDatabaseURL := params.MigrationsDatabaseURL
	migrationSourceURL := params.MigrationsSourceURL

	migrations, err := migrate.New(migrationSourceURL, migrationDatabaseURL)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf(
			"error creating migration: %v", err)
	}

	if err := migrations.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println(err)
		return fmt.Errorf("error applying migrations: %v", err)
	}

	fmt.Println("Migrations successfully applied")
	return nil
}
