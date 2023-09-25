package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"

	// Package
	api "mumogo/controller/api/stageGreeting"
	"mumogo/controller/auth"
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
	authController          = auth.NewAuthController()
	stageGreetingController = api.NewStageGreetingController()
)

func setupRouter() *gin.Engine {
	router := gin.Default()

	// CORS 설정
	config := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001", "http://localhost:8080", "https://port-0-mumo-1xxfe2bllyrqhy9.sel5.cloudtype.app", "https://web-mumo-react-1xxfe2bllyrqhy9.sel5.cloudtype.app"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	router.Use(config)
	// router.Use(authController.CheckAccessToken)

	// 라우팅 설정
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition

	// "api/open/*" 패턴에 대한 JWT 토큰 체크를 하지 않음
	// openGroup := router.Group("/api/open")
	// {
	// 	openGroup.GET("/stageGreetings", stageGreetingController.GetStageGreetingUrls)
	// }

	router.GET("/api/v1/stageGreetingUrls", stageGreetingController.GetStageGreetingUrls)
	router.GET("/api/v1/stageGreetings", stageGreetingController.GetStageGreetings)

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

	go func() {
		// gocron 스케줄러 생성
		s := gocron.NewScheduler(time.UTC)

		// 5분 간격으로 실행될 작업을 정의
		// _, err = s.Every(5).Minutes().Do(crawler.NewCrawlerController().CrawlMegabox)

		if err != nil {
			fmt.Printf("NewCrawlerController Error: %s\n", err)
			return
		}

		// 스케줄러 시작
		s.StartBlocking()
	}()

	// crawler.NewCrawlerController().CrawlMegabox()
	// crawler.NewCrawlerController().CrawlLotteCinema()
	// crawler.NewCrawlerController().CrawlCgv()
	// crawler.NewCrawlerController().CrawlTest()

	// userRepository := repository.NewUserRepository()
	// users, err := userRepository.GetUsers()
	// if err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user"})
	// 	print(err)
	// 	// return
	// }

	// print(users)

	router.Run(":8080")
	// router.Run("0.0.0.0:8080")

}
