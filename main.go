package main

import (
	"github.com/gin-gonic/gin"

	// Package
	"mumogo/controller/album"
	"mumogo/controller/auth"
	"mumogo/controller/crawler"
	"mumogo/controller/stageGreeting"
	"mumogo/model"

	// Security
	cors "github.com/rs/cors/wrapper/gin"

	// DB
	"mumogo/db"

	// Swagger
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

var (
	authController = auth.NewAuthController()
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS 설정
	config := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(config)
	router.Use(authController.CheckAccessToken)

	// 라우팅 설정
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	router.GET("/test", stageGreeting.GetStageGreetings)
	router.GET("/albums", album.GetAlbums)
	router.POST("/auth/login", authController.Login)
	router.GET("/auth/users", authController.GetUsers)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {

	router := setupRouter()

	err := db.InitDB()
	if err != nil {
		// 데이터베이스 연결 초기화 실패 처리
		return
	}

	db.DB.AutoMigrate(&model.User{})
	db.DB.AutoMigrate(&model.Movie{})
	db.DB.AutoMigrate(&model.StageGreeting{})
	db.DB.AutoMigrate(&model.StageGreetingUrl{})

	crawler.NewCrawlerController().CrawlMegabox()
	// crawler.NewCrawlerController().CrawlLotteCinema()
	crawler.NewCrawlerController().CrawlCgv()
	// crawler.NewCrawlerController().CrawlTest()

	// userRepository := repository.NewUserRepository()
	// users, err := userRepository.GetUsers()
	// if err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
	// 	print(err)
	// 	// return
	// }

	// print(users)

	router.Run("localhost:8080")
}