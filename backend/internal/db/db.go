package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DbClient struct {
	DB *gorm.DB
}

func NewDbClient() (*DbClient, error) {
	db, err := gorm.Open(sqlite.Open("notes.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbClient := DbClient{
		DB: db,
	}
	return &dbClient, dbClient.migrate()
}

func (d *DbClient) migrate() error {
	return d.DB.AutoMigrate(
		&Note{},
	)
}
