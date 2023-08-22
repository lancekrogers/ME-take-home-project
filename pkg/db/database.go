package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/pkg/errors"
)

type DatabaseSettings struct {
	Host           string
	Port           string
	User           string
	Password       string
	Database       string
	SSLModeDisable bool
}

func GetDatabase(dbSettings DatabaseSettings) (*sql.DB, error) {
	connStr, err := GetConnectionString(dbSettings)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s", connStr))
	if err != nil {
		return nil, errors.Wrap(err, "error opening the database")
	}
	err = PingDB(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to verify connection to database at hostname:port %s:%s", dbSettings.Host, dbSettings.Port))
	}

	log.Println("Database connection successful")
	return db, nil
}

func GetConnectionString(dbSettings DatabaseSettings) (string, error) {
	conStr := fmt.Sprintf("%s:%s@%s:%s/%s?sslmode=disable", dbSettings.User, dbSettings.Password, dbSettings.Host, dbSettings.Port, dbSettings.Database)
	return conStr, nil
}

func PingDB(db *sql.DB) error {
	deadline := time.Now().Add(60 * time.Second)
	err := errors.New("something went wrong")
	for time.Now().Before(deadline) {
		err = db.Ping()
		if err == nil {
			return nil
		}
	}
	if err != nil {
		return errors.Wrap(err, "failed to ping database")
	}
	return nil
}

func MigrateAndGetDatabaseWithIOFS(source source.Driver, dbSettings DatabaseSettings) (*sql.DB, error) {
	connectionString, err := GetConnectionString(dbSettings)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create connection string")
	}
	uri := "postgres://" + connectionString
	log.Println("Migrating database schema")
	m, err := migrate.NewWithSourceInstance("iofs", source, uri)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to initialize migrations")
	}
	if err := m.Up(); err != nil && err.Error() != "no change" {
		return nil, errors.Wrap(err, "error migrating database schema")
	}
	log.Println("Schema up to date")

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, errors.Wrap(err, "error opening the database")
	}
	err = PingDB(db)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("failed to verify connectivity of database at hostname:port %s:%s", dbSettings.Host, dbSettings.Port))
	}

	log.Println("Database connection successful")
	return db, nil
}
