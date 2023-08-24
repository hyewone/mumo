package service

import (
	"mumogo/model"
	"mumogo/repository"
)

type CrawlerService struct {
	Repo *repository.CrawlerRepository
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{
		Repo: repository.NewCrawlerRepository(),
	}
}

func (s *CrawlerService) AddMovie(movies []*model.Movie) error {
	for _, movie := range movies {
		err := s.Repo.AddMovie(movie)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *CrawlerService) GetMovieByName(name string) (*model.Movie, error) {
	return s.Repo.GetMovieByName(name)
}

func (s *CrawlerService) AddStageGreeting(sgs []*model.StageGreeting) error {
	for _, sg := range sgs {
		err := s.Repo.AddStageGreeting(sg)
		if err != nil {
			return err
		}
	}
	return nil
}
func (s *CrawlerService) AddStageGreetingUrl(urls []*model.StageGreetingUrl) error {
	for _, url := range urls {
		err := s.Repo.AddStageGreetingUrl(url)
		if err != nil {
			return err
		}
	}
	return nil
}

// func (s *UserService) IsUserExists(email string, provider string) bool {
// 	_, err := s.Repo.GetUserByEmailAndProvider(email, provider)
// 	return err == nil
// }

// func (s *UserService) GetUserByEmailAndProvider(email string, provider string) (*model.User, error) {
// 	return s.Repo.GetUserByEmailAndProvider(email, provider)
// }

// func (s *UserService) GetUserById(id int) (*model.User, error) {
// 	return s.Repo.GetUserById(id)
// }

// func (s *UserService) CreateUser(email string, provider string) error {
// 	user := &model.User{
// 		Email:    email,
// 		Provider: provider,
// 		UserType: "USER",
// 	}
// 	return s.Repo.CreateUser(user)
// }

// func (s *UserService) GetUsers() ([]model.User, error) {
// 	return s.Repo.GetUsers()
// }
