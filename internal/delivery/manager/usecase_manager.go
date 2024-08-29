package manager

import "golang-gorm/internal/usecase"

type UsecaseManager interface {
	UserManager() usecase.UsecaseUser
}

type usecaseManager struct {
	repo RepoManager
}

func (uc usecaseManager) UserManager() usecase.UsecaseUser {
	return usecase.NewUsecaseUser(uc.repo.UserRepo(), uc.repo.RoleUserRepo())
}

func NewUsecaseManager(repo RepoManager) UsecaseManager {
	return &usecaseManager{repo: repo}
}
