package user

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"dashboard/service"
)

type Controller struct{}

func (ctrl Controller) Index(c *gin.Context) {
	var s user.Service
	result, err := s.GetAll()

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl Controller) Create(c *gin.Context) {
	var s user.Service
	result, err := s.CreateModel(c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(201, result)
	}
}

func (ctrl Controller) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	result, err := s.GetByID(id)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl Controller) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	result, err := s.UpdateByID(id, c)

	if err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(200, result)
	}
}

func (ctrl Controller) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var s user.Service
	if err := s.DeleteByID(id); err != nil {
		c.AbortWithStatus(500)
		fmt.Println(err)
	} else {
		c.JSON(204, gin.H{"id #" + id: "deleted"})
	}
}
