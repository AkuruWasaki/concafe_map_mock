package controllers

import (
	"net/http"
	"strconv"

	"github.com/akuruwasaki/concafe_map_mock/db"
	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Controller is shops controller
type ShopsController struct{}

type Shop struct {
	Shop       models.Shop `json:"shop"`
	ShopGenres []int       `json:"shop_genre_ids"`
}

// Index action: GET /shops
func (sc ShopsController) Index(c *gin.Context) {
	db := db.Connect()
	shops, err := models.Shops(qm.Limit(5)).All(c, db)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, shops)
}

// Create action: POST /shops
func (sc ShopsController) Create(c *gin.Context) {
	db := db.Connect()
	var shop Shop

	// set param
	if err := c.BindJSON(&shop); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := shop.Shop.Insert(c, db, boil.Infer()); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, genre_id := range shop.ShopGenres {
		var shop_genre_relation models.ShopGenreRelation
		shop_genre_relation.ShopGenreID = genre_id
		shop_genre_relation.ShopID = shop.Shop.ID
		// insert
		if err := shop_genre_relation.Insert(c, db, boil.Infer()); err != nil {
			c.AbortWithStatus(400)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(201, shop)
}

// Update action: PUT /shops/:id
func (sc ShopsController) Update(c *gin.Context) {
	db := db.Connect()

	// パスからID取得
	id := c.Param("id")
	// idをint型に変換
	idInt, _ := strconv.Atoi(id)

	// find shop data
	shop, err := models.FindShop(c, db, idInt)
	// ToDo: c.AbortWithStatus(400)の部分をうまいこと共通化したい
	if shop == nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop not found"})
		return
	}

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set param
	c.BindJSON(&shop)

	_, err = shop.Update(c, db, boil.Infer())
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "の店舗情報を更新しました"})
}

// Delete action: DELETE /shops/:id
func (sc ShopsController) Delete(c *gin.Context) {
	db := db.Connect()

	// パスからID取得
	id := c.Param("id")
	// idをint型に変換
	idInt, _ := strconv.Atoi(id)

	// find shop data
	shop, err := models.FindShop(c, db, idInt)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// delete shop
	_, err = shop.Delete(c, db)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "の店舗情報を削除しました"})
}
