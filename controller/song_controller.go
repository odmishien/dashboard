package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dashboard/service"
)

type SongController struct{}

func (ctrl SongController) FindByName(c *gin.Context) {
	name := c.Query("name")
	var s service.SongService
	result, err := s.FindByName(name)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}