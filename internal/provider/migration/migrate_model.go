package migration

import (
	"golang-gorm/internal/domain"
	"gorm.io/gorm"
)

type ModelMigration interface {
	Migrate(db *gorm.DB) error
}

type modelMigration struct {
	db *gorm.DB
}

func (m modelMigration) Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		return err
	}

	return nil
}

func NewModelMigration(db *gorm.DB) ModelMigration {
	return &modelMigration{db: db}
}
