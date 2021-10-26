package server

import (
	"github.com/akuruwasaki/concafe_map_mock/controllers"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init() {
	r := GetRouter()
	r.Run()
}

func GetRouter() *gin.Engine {
	r := gin.Default()

	s := r.Group("/shops")
	{
		ctrl := controllers.ShopsController{}
		s.GET("", ctrl.Index)
		s.POST("", ctrl.Create)
		s.PUT("/:id", ctrl.Update)
		s.DELETE("/:id", ctrl.Delete)
	}

	return r
}
