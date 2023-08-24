package repository

import (
	"mumogo/db"
	"mumogo/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: db.GetDB(),
	}
}

func (r *UserRepository) GetUserByEmailAndProvider(email string, provider string) (*model.User, error) {
	var user model.User
	err := r.DB.Where("email = ? AND provider = ?", email, provider).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetUserById(id int) (*model.User, error) {
	var user model.User
	err := r.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *model.User) error {
	return r.DB.Create(user).Error
}

func (r *UserRepository) GetUsers() ([]model.User, error) {
	var users []model.User
	err := r.DB.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
