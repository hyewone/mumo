package repository

import (
	"mumogo/db"
	"mumogo/model"

	"gorm.io/gorm"
)

type StageGreetingRepository struct {
	DB *gorm.DB
}

func NewStageGreetingRepository() *StageGreetingRepository {
	return &StageGreetingRepository{
		DB: db.GetDB(),
	}
}

func (r *StageGreetingRepository) GetStageGreetingUrls(cinemaType string) ([]model.StageGreetingUrl, error) {
	var stageGreetingUrls []model.StageGreetingUrl
	err := r.DB.Find(&stageGreetingUrls, "cinema_type = ?", cinemaType).Error
	if err != nil {
		return nil, err
	}
	return stageGreetingUrls, nil
}

// func (r *UserRepository) GetUserByEmailAndProvider(email string, provider string) (*model.User, error) {
// 	var user model.User
// 	err := r.DB.Where("email = ? AND provider = ?", email, provider).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *UserRepository) GetUserById(id int) (*model.User, error) {
// 	var user model.User
// 	err := r.DB.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *UserRepository) CreateUser(user *model.User) error {
// 	return r.DB.Create(user).Error
// }

// func (r *UserRepository) GetUsers() ([]model.User, error) {
// 	var users []model.User
// 	err := r.DB.Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
