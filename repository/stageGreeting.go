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
	// var stageGreetingUrlsWithMovieName []model.StageGreetingUrlWithMovieName
	// err := r.DB.Find(&stageGreetingUrls, "cinema_type = ?", cinemaType).Error

	err := r.DB.Model(&model.StageGreetingUrl{}).Preload("Movie").Find(&stageGreetingUrls, "cinema_type = ?", cinemaType).Error

	// err := r.DB.Table("stage_greeting_urls").
	// 	Select("stage_greeting_urls.*, movies.name").
	// 	Joins("INNER JOIN movies ON movies.id = stage_greeting_urls.movie_id").
	// 	Find(&stageGreetingUrls).
	// 	Error

	// 	db.Joins("Company").Find(&users)
	// // SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	// // inner join
	// db.InnerJoins("Company").Find(&users)
	// // SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` INNER JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;

	// err := r.DB.
	// 	Table("stage_greeting_urls").
	// 	Select("stage_greeting_urls.*, movies.name as MovieNm").
	// 	Joins("JOIN movies ON stage_greeting_urls.movie_id = movies.id").
	// 	Where("stage_greeting_urls.cinema_type = ?", cinemaType).
	// 	Scan(&stageGreetingUrlsWithMovieName).
	// 	Error
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
