package service

import (
	"mumogo/model"
	"mumogo/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		Repo: repository.NewUserRepository(),
	}
}

func (s *UserService) IsUserExists(email string, provider string) bool {
	_, err := s.Repo.GetUserByEmailAndProvider(email, provider)
	return err == nil
}

func (s *UserService) GetUserByEmailAndProvider(email string, provider string) (*model.User, error) {
	return s.Repo.GetUserByEmailAndProvider(email, provider)
}

func (s *UserService) GetUserById(id int) (*model.User, error) {
	return s.Repo.GetUserById(id)
}

func (s *UserService) CreateUser(email string, provider string) error {
	user := &model.User{
		Email:    email,
		Provider: provider,
		UserType: "USER",
	}
	return s.Repo.CreateUser(user)
}

func (s *UserService) GetUsers() ([]model.User, error) {
	return s.Repo.GetUsers()
}
