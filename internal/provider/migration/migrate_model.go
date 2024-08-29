package migration

import (
	"golang-gorm/internal/domain"
	"gorm.io/gorm"
)

// ModelMigration interface yang memiliki metode Migrate untuk migrasi database
type ModelMigration interface {
	Migrate() error
	InputData() error
}

// modelMigration struct yang menyimpan instance dari *gorm.DB
type modelMigration struct {
	db *gorm.DB
}

func (m *modelMigration) InputData() error {

	roles := []*domain.RoleUser{
		{
			ID:       "ROLE-001",
			RoleName: "ADMIN",
		},
		{
			ID:       "ROLE-002",
			RoleName: "USER",
		},
	}

	result := m.db.Create(roles)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// Migrate melakukan migrasi dengan menggunakan GORM migrator
func (m *modelMigration) Migrate() error {
	migrate := m.db.Migrator()

	//Drop tabel yang ada
	//err := migrate.DropTable(
	//	&domain.User{},
	//	&domain.RoleUser{},
	//	&domain.Product{},
	//	&domain.Payment{},
	//	&domain.Order{},
	//	&domain.OrderItem{},
	//	&domain.Category_Product{},
	//)
	//if err != nil {
	//	return err
	//}

	// AutoMigrate untuk membuat tabel sesuai dengan model domain
	err := migrate.AutoMigrate(
		&domain.User{},
		&domain.RoleUser{},
		&domain.Product{},
		&domain.Payment{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Category_Product{},
	)
	if err != nil {
		return err
	}

	return nil
}

// NewModelMigration membuat instance baru dari ModelMigration
func NewModelMigration(db *gorm.DB) (ModelMigration, error) {
	return &modelMigration{db: db}, nil
}
