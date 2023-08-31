package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// var sqlDB *sql.DB
// var gormDB *gorm.DB
var DB *gorm.DB

func InitDB() error {
	// dsn := "host=localhost user=postgre_local password=1234 dbname=mumo port=5432 sslmode=disable TimeZone=Asia/Seoul"
	dsn := "host=database.c0u6xj9tqa4x.ap-northeast-2.rds.amazonaws.com user=hyenee password=skfkrh12 dbname=mumo port=5432 sslmode=disable TimeZone=Asia/Seoul"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return err
	}

	return nil
}

func GetDB() *gorm.DB {
	if DB == nil {
		err := InitDB()
		if err != nil {
			return nil
		}
	}
	return DB
}
