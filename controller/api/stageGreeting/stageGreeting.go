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

// router.GET("/open/megabox/stageGreetings", stageGreetingController.GetMegaboxStageGreetings)
// router.GET("/open/lottecinema/stageGreetings", stageGreetingController.GetLottecinemaStageGreetings)
// router.GET("/open/cgv/stageGreetings", stageGreetingController.GetCgvStageGreetings)

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
