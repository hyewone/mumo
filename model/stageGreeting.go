package model

import (
	"gorm.io/gorm"
)

type StageGreeting struct {
	gorm.Model // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt)

	// MovieID        int    // 영화ID
	Movie          Movie `gorm:"foreignKey:MovieID"`
	MovieID        int
	CinemaType     string // MEGABOX, LOTTECINEMA, CGV
	Theater        string // 극장명
	ShowDate       string // 상영날짜
	ShowTime       string // 상영시간
	RemainingSeats int    // 잔여좌석
	AttendeeName   string // 참석자
}

func NewStageGreeting(movie Movie, cinemaType string, theater string, showDate string, showTime string, remainingSeats int, attendeeName string) *StageGreeting {
	return &StageGreeting{
		// MovieID:        movieId,
		Movie:          movie,
		CinemaType:     cinemaType,
		Theater:        theater,
		ShowDate:       showDate,
		ShowTime:       showTime,
		RemainingSeats: remainingSeats,
		AttendeeName:   attendeeName,
	}
}
