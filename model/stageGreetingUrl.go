package model

import (
	"gorm.io/gorm"
)

type StageGreetingUrl struct {
	gorm.Model // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt)

	// MovieID    int    // 영화ID
	Movie      Movie `gorm:"foreignKey:MovieID"`
	MovieID    int
	CinemaType string // MEGABOX, LOTTECINEMA, CGV
	Title      string // 게시글 제목
	Url        string // 시네마별 무대인사 Url
	Image      string // 게시글 이미지
	EndYn      string // 종료 여부?
}

func NewStageGreetingUrl(movie Movie, cinemaType string, title string, url string, img string, endYn string) *StageGreetingUrl {
	return &StageGreetingUrl{
		// MovieID:    movieId,
		Movie:      movie,
		CinemaType: cinemaType,
		Title:      title,
		Url:        url,
		Image:      img,
		EndYn:      endYn,
	}
}
