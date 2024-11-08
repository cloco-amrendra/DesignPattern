package usecase

import (
	"go-design-patterns/singleton/models"
	"go-design-patterns/singleton/repository"
)

type UserUsecase struct {
	userRepo *repository.UserRepository
}

func NewUserUsecase(userRepo *repository.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: userRepo}
}

func (s *UserUsecase) GetUsers() ([]models.User, error) {
	return s.userRepo.GetAllUsers()
}

func (s *UserUsecase) CreateUser(user *models.User) error {
	return s.userRepo.CreateUser(user)
}
