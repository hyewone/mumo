package repository

import (
	"mumogo/db"
	"mumogo/model"

	"gorm.io/gorm"
)

type CrawlerRepository struct {
	DB *gorm.DB
}

func NewCrawlerRepository() *CrawlerRepository {
	return &CrawlerRepository{
		DB: db.GetDB(),
	}
}

func (r *CrawlerRepository) AddMovie(movie *model.Movie) error {
	var existingMovie model.Movie

	err := r.DB.Where(model.Movie{Name: movie.Name}).FirstOrCreate(&existingMovie).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *CrawlerRepository) GetMovieByName(name string) (*model.Movie, error) {
	var movie model.Movie
	err := r.DB.Where("name = ?", name).First(&movie).Error
	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (r *CrawlerRepository) AddStageGreeting(stageGreeting *model.StageGreeting) error {
	var existingSg model.StageGreeting

	keyAttrs := model.StageGreeting{
		MovieID:    stageGreeting.MovieID,
		CinemaType: stageGreeting.CinemaType,
		Theater:    stageGreeting.Theater,
		ShowDate:   stageGreeting.ShowDate,
		ShowTime:   stageGreeting.ShowTime,
	}

	valueAttrs := model.StageGreeting{
		RemainingSeats: stageGreeting.RemainingSeats,
		AttendeeName:   stageGreeting.AttendeeName,
	}

	err := r.DB.Where(keyAttrs).Assign(valueAttrs).FirstOrCreate(&existingSg).Error
	if err != nil {
		return err
	}

	return nil
}
func (r *CrawlerRepository) AddStageGreetingUrl(stageGreetingUrl *model.StageGreetingUrl) error {
	var existingUrls model.StageGreetingUrl

	keyAttrs := model.StageGreetingUrl{
		MovieID:    stageGreetingUrl.MovieID,
		CinemaType: stageGreetingUrl.CinemaType,
		Title:      stageGreetingUrl.Title,
	}

	valueAttrs := model.StageGreetingUrl{
		Image: stageGreetingUrl.Image,
		Url:   stageGreetingUrl.Url,
		EndYn: stageGreetingUrl.EndYn,
	}

	err := r.DB.Where(keyAttrs).Assign(valueAttrs).FirstOrCreate(&existingUrls).Error
	if err != nil {
		return err
	}

	return nil
}

// func (r *CrawlerRepository) GetUserByEmailAndProvider(email string, provider string) (*model.User, error) {
// 	var user model.User
// 	err := r.DB.Where("email = ? AND provider = ?", email, provider).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *CrawlerRepository) GetUserById(id int) (*model.User, error) {
// 	var user model.User
// 	err := r.DB.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (r *CrawlerRepository) CreateUser(user *model.User) error {
// 	return r.DB.Create(user).Error
// }

// func (r *CrawlerRepository) GetUsers() ([]model.User, error) {

// 	var users []model.User
// 	err := r.DB.Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return users, nil
// }
