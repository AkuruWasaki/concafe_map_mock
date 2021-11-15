package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/akuruwasaki/concafe_map_mock/models/modext"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Controller is shops controller
type ShopsController struct {
	DB *sql.DB
}

type Shop struct {
	Shop       models.Shop `json:"shop"`
	ShopGenres []int       `json:"shop_genre_ids"`
}

// Index action: GET /shops
func (sc ShopsController) Index(c *gin.Context) {
	shops, err := models.Shops(qm.Limit(5)).All(c, sc.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, shops)
}

// Create action: POST /shops
func (sc ShopsController) Create(c *gin.Context) {
	var shop Shop

	// set param
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validation
	if errMsgs := modext.ValidateShop(&shop.Shop); errMsgs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsgs})
		return
	}

	// start transaction
	tx, err := sc.DB.BeginTx(c, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := shop.Shop.Insert(c, tx, boil.Infer()); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 店舗のジャンルは最大5店舗を想定しているため, for rangeで対応
	for _, genre_id := range shop.ShopGenres {
		var shop_genre_relation models.ShopGenreRelation
		shop_genre_relation.ShopGenreID = genre_id
		shop_genre_relation.ShopID = shop.Shop.ID
		// insert
		if err := shop_genre_relation.Insert(c, tx, boil.Infer()); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	c.JSON(201, shop)
}

// Update action: PUT /shops/:id
func (sc ShopsController) Update(c *gin.Context) {
	// idをint型に変換
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find shop data
	shop, err := models.FindShop(c, sc.DB, idInt)
	if shop == nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "shop not found"})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set param
	if err := c.BindJSON(&shop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 店舗情報アップデート
	_, err = shop.Update(c, sc.DB, boil.Infer())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + strconv.Itoa(idInt) + "の店舗情報を更新しました"})
}

// Delete action: DELETE /shops/:id
func (sc ShopsController) Delete(c *gin.Context) {
	// パスからID取得
	// idをint型に変換
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find shop data
	shop, err := models.FindShop(c, sc.DB, idInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// start transaction
	tx, err := sc.DB.BeginTx(c, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 関連テーブルのデータを削除する。
	// 店舗ジャンル削除
	if _, err := models.
		ShopGenreRelations(models.ShopGenreRelationWhere.ShopID.EQ(idInt)).
		DeleteAll(c, tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// 店舗スタッフ削除
	if _, err := models.Staffs(models.StaffWhere.ShopID.EQ(idInt)).
		DeleteAll(c, tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}

	// delete shop
	if _, err = shop.Delete(c, tx); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		tx.Rollback()
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + strconv.Itoa(idInt) + "の店舗情報を削除しました"})
	tx.Commit()
}
