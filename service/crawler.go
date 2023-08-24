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
