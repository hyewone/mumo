package stageGreeting

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

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
