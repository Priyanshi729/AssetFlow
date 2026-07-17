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

func ConnectDB(host, port, user, password, databaseName string) error {
	var err error

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)

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

func Tx(fn func(tx *sqlx.Tx) error) (err error) {
	tx, err := DB.Beginx()
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Printf("failed to rollback transaction: %v", rollbackErr)
			}
			return
		}

		if commitErr := tx.Commit(); commitErr != nil {
			err = fmt.Errorf("failed to commit transaction: %w", commitErr)
		}
	}()
	err = fn(tx)
	return
}
