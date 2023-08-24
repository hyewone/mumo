package model

import (
	"gorm.io/gorm"
)

type StageGreetingUrl struct {
	gorm.Model // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt)

	MovieID    int    // 영화ID
	CinemaType string // MEGABOX, LOTTECINEMA, CGV
	Title      string // 게시글 제목
	Url        string // 시네마별 무대인사 Url
	Image      string // 게시글 이미지
	EndYn      string // 종료 여부?
}

func NewStageGreetingUrl(movieId int, cinemaType string, title string, url string, img string, endYn string) *StageGreetingUrl {
	return &StageGreetingUrl{
		MovieID:    movieId,
		CinemaType: cinemaType,
		Title:      title,
		Url:        url,
		Image:      img,
		EndYn:      endYn,
	}
}
