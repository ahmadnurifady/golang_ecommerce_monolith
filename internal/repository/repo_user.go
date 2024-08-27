package repository

import (
	"fmt"
	"golang-gorm/internal/domain"
	"gorm.io/gorm"
)

type RepositoryUser interface {
	CreateUser(payload domain.User) (string, error)
	FindUserById(id string) (domain.User, error)
}

type repositoryUser struct {
	db *gorm.DB
}

func (repo repositoryUser) CreateUser(payload domain.User) (string, error) {
	result := repo.db.Create(&payload)

	if result.RowsAffected != 1 {
		return "", fmt.Errorf("failed to create the user")
	}

	return fmt.Sprintf("success create user with ID = %d", payload.ID), nil

}

func (repo repositoryUser) FindUserById(id string) (domain.User, error) {
	var user domain.User
	result := repo.db.First(&user, id)
	if result.RowsAffected != 1 {
		return domain.User{}, fmt.Errorf("failed to find the user")
	}
	return user, nil
}

func NewRepositoryUser(db *gorm.DB) RepositoryUser {
	return &repositoryUser{db: db}
}
