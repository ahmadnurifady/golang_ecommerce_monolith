package usecase

import (
	"github.com/google/uuid"
	"golang-gorm/internal/domain"
	"golang-gorm/internal/domain/dto"
	"golang-gorm/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UsecaseUser interface {
	CreateUserUsecase(payload dto.UserRequest) (string, error)
	FindUserByIdUsecase(id string) (domain.User, error)
}

type usecaseUser struct {
	repo repository.RepositoryUser
}

func (uc usecaseUser) CreateUserUsecase(payload dto.UserRequest) (string, error) {

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	dataToCreate := domain.User{
		ID:        uuid.New().String(),
		Username:  payload.Username,
		Email:     payload.Email,
		Password:  string(hashPassword),
		Role:      "ADMIN",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	result, err := uc.repo.CreateUser(dataToCreate)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (uc usecaseUser) FindUserByIdUsecase(id string) (domain.User, error) {
	result, err := uc.repo.FindUserById(id)
	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}

func NewUsecaseUser(repo repository.RepositoryUser) UsecaseUser {
	return &usecaseUser{repo: repo}
}
