package config

import (
	"errors"
	"os"

	"github.com/enyasantos/project-manager/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_CONNECTION")

	if dsn == "" {
		return nil, errors.New("The DATABASE_DNS environment variable is not defined")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if pgDB, err := db.DB(); err != nil {
		return nil, err
	} else {
		err = pgDB.Ping()
		if err != nil {
			return nil, err
		}
	}

	err = db.AutoMigrate(&schemas.Project{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	logger.Info("Postgres is running successfully.")

	return db, nil
}
