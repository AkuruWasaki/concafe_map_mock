package controllers

import (
	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"net/http"
	"time"
)

// Controller is shops controller
type ShopsController struct{}

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

// Create action: POST /shops
func (sc ShopsController) Create(c *gin.Context) {
	db := db.Connect()
	var shop models.Shop

	// WIP: set param
	shop.Name = c.Param("name")
	shop.Address = null.StringFrom(c.Param("address"))
	shop.Tel = null.StringFrom(c.Param("tel"))
	shop.Content = null.StringFrom(c.Param("content"))
	shop.CreatedAt = time.Now()
	shop.UpdatedAt = time.Now()

	if err := shop.Insert(c, db, boil.Infer()); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
