package database

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sqlx.DB

func ConnectDB(host, port, user, password, databaseName, sslmode string) error {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, databaseName, sslmode)

	DB, err = sqlx.Connect("postgres", dsn)

	if err != nil {
		return err
	}
	log.Println("Connected to DB")

	return migrateUp(DB)

}

func migrateUp(db *sqlx.DB) error {
	driver, driErr := postgres.WithInstance(db.DB, &postgres.Config{})

	if driErr != nil {
		return driErr
	}
	m, migErr := migrate.NewWithDatabaseInstance("file://database/migration", "postgres", driver)

	if migErr != nil {
		return migErr
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func CloseDB() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
