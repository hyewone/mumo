package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model // GORM의 기본 모델 (ID, CreatedAt, UpdatedAt, DeletedAt)

	Email    string `gorm:"uniqueIndex"` // 이메일, 고유 인덱스 설정
	Provider string // GOOGLE, KAKAO, NAVER
	UserType string // USER, ADMIN
}

func NewUser(email, provider, userType string) *User {
	if userType == "" || userType == "null" {
		userType = "USER"
	}
	return &User{
		Email:    email,
		Provider: provider,
		UserType: userType,
	}
}
