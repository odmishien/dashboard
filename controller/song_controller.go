package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dashboard/service"
)

type SongController struct{}

func (ctrl SongController) Index(c *gin.Context) {
	var s service.SongService
	result, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl SongController) Create(c *gin.Context) {
	var s service.SongService
	result, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, result)
	}
}

func (ctrl SongController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.SongService
	result, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl SongController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.SongService
	result, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl SongController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.SongService
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
