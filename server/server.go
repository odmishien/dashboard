package server

import (
	"github.com/gin-gonic/gin"

	"dashboard/controller"
)

// Init is initialize server
func Init() {
	r := router()
	r.Run()
}

func router() *gin.Engine {
	r := gin.Default()

	u := r.Group("/users")
	{
		ctrl := controller.UserController{}
		u.GET("", ctrl.Index)
		u.GET("/:id", ctrl.Show)
		u.POST("", ctrl.Create)
		u.PUT("/:id", ctrl.Update)
		u.DELETE("/:id", ctrl.Delete)
	}

	ws := r.Group("/wishsongs")
	{
		ctrl := controller.WishSongController{}
		ws.GET("", ctrl.Index)
		ws.GET("/:id", ctrl.Show)
		ws.POST("", ctrl.Create)
		ws.PUT("/:id", ctrl.Update)
		ws.DELETE("/:id", ctrl.Delete)
	}

	s := r.Group("/songs")
	{
		ctrl := controller.SongController{}
		s.GET("", ctrl.FindByName)
	}

	return r
}
