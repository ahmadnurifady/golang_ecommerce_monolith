package manager

import (
	"golang-gorm/internal/provider/db"
	"golang-gorm/internal/repository"
)

type RepoManager interface {
	UserRepo() repository.RepositoryUser
	RoleUserRepo() repository.RepoRoleUser
}

type repoManager struct {
	db db.ConnectDatabase
}

func (r repoManager) UserRepo() repository.RepositoryUser {
	return repository.NewRepositoryUser(r.db.Conn())
}

func (r repoManager) RoleUserRepo() repository.RepoRoleUser {
	return repository.NewRepoRoleUser(r.db.Conn())
}

func NewRepoManager(db db.ConnectDatabase) RepoManager {
	return &repoManager{db: db}
}
