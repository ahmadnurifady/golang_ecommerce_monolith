package repository

import (
	"fmt"
	"golang-gorm/internal/domain"
	"gorm.io/gorm"
)

type RepoRoleUser interface {
	FindRoleByName(roleName string) (domain.RoleUser, error)
}

type repoRoleUser struct {
	db *gorm.DB
}

func (repo repoRoleUser) FindRoleByName(roleName string) (domain.RoleUser, error) {
	var role domain.RoleUser
	result := repo.db.First(&role, "role_name", roleName)
	if result.RowsAffected != 1 {
		return domain.RoleUser{}, fmt.Errorf("failed to find role user")
	}
	return role, nil
}

func NewRepoRoleUser(db *gorm.DB) RepoRoleUser {
	return &repoRoleUser{
		db: db,
	}
}
