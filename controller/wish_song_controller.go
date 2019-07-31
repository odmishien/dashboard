package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dashboard/service"
)

type WishSongController struct{}

func (ctrl WishSongController) Index(c *gin.Context) {
	var s service.WishSongService
	result, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl WishSongController) Create(c *gin.Context) {
	var s service.WishSongService
	result, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, result)
	}
}

func (ctrl WishSongController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.WishSongService
	result, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl WishSongController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.WishSongService
	result, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl WishSongController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s service.WishSongService
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
