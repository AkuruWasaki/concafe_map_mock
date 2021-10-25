package controllers

import (
	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
)

// Controller is shops controller
type ShopsController struct{}

// WIP
// shops_controller
// 店舗一覧
// func ShopIndex(c *gin.Context) {
// 	db := db.Connect()
// 	shops := models.Shops()
// 	c.JSON(200, shops)
// }

// Index action: GET /shops
func (sc ShopsController) Index(c *gin.Context) {
	db := db.Connect()
	shops, err := models.Shops(qm.Limit(5)).All(c, db)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, shops)
	}
}
