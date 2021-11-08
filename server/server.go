package server

import (
	"database/sql"

	"github.com/akuruwasaki/concafe_map_mock/controllers"
	"github.com/gin-gonic/gin"
)

// Init is initialize server
func Init(db *sql.DB) {
	r := GetRouter(db)
	r.Run()
}

func GetRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	s := r.Group("/shops")
	{
		ctrl := controllers.ShopsController{}
		ctrl.DB = db
		s.GET("", ctrl.Index)
		s.POST("", ctrl.Create)
		s.PUT("/:id", ctrl.Update)
		s.DELETE("/:id", ctrl.Delete)
	}

	stf := r.Group("/staffs")
	{
		ctrl := controllers.StaffsController{}
		ctrl.DB = db
		stf.GET("", ctrl.Index)
		stf.POST("", ctrl.Create)
		stf.PUT("/:id", ctrl.Update)
		stf.DELETE("/:id", ctrl.Delete)
	}

	return r
}
