package api

import (
	"mumogo/service"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type StageGreetingController struct {
	Service *service.StageGreetingService
}

func NewStageGreetingController() *StageGreetingController {
	return &StageGreetingController{
		Service: service.NewStageGreetingService(),
	}
}

func (con *StageGreetingController) GetStageGreetingUrls(c *gin.Context) {
	cinemaType := c.Query("cinemaType")
	cinemaType = strings.ToUpper(cinemaType)

	stageGreetingUrls, err := con.Service.GetStageGreetingUrls(cinemaType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while GetStageGreetingUrls"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stageGreetingUrls": stageGreetingUrls})
}

func (con *StageGreetingController) GetStageGreetings(c *gin.Context) {
	// cinemaType := c.Query("cinemaType")
	// cinemaType = strings.ToUpper(cinemaType)

	stageGreetings, err := con.Service.GetStageGreetings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while GetStageGreetings"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"stageGreetings": stageGreetings})
}

// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param userBody body User true "User Info Body"
// @Success 200 {object} User
// @Router /user [post]
func GetStageGreetings(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func GetStageGreetingById(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}

func PostStageGreeting(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, nil)
}
