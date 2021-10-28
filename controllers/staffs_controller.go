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

// Controller is staffs controller
type StaffsController struct{}

// Index action: GET /staffs
func (sc StaffsController) Index(c *gin.Context) {
	db := db.Connect()
	staffs, err := models.Staffs(qm.Limit(5)).All(c, db)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, staffs)
	}
}

// Create action: POST /staffs
func (sc StaffsController) Create(c *gin.Context) {
	db := db.Connect()
	var staff models.Staff

	c.BindJSON(&staff)

	if err := staff.Insert(c, db, boil.Infer()); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, staff)
}

// Update action: PUT /staffs/:id
func (sc StaffsController) Update(c *gin.Context) {
	db := db.Connect()

	// パスからID取得
	id := c.Param("id")
	// idをint型に変換
	idInt, _ := strconv.Atoi(id)

	// find staff data
	staff, err := models.FindStaff(c, db, idInt)
	// ToDo: c.AbortWithStatus(400)の部分をうまいこと共通化したい
	if staff == nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// set param
	c.BindJSON(&staff)

	_, err = staff.Update(c, db, boil.Infer())
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "のスタッフ情報を更新しました"})
}

// Delete action: DELETE /staffs/:id
func (sc StaffsController) Delete(c *gin.Context) {
	db := db.Connect()

	// パスからID取得
	id := c.Param("id")
	// idをint型に変換
	idInt, _ := strconv.Atoi(id)

	// find staff data
	staff, err := models.FindStaff(c, db, idInt)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// delete staff
	_, err = staff.Delete(c, db)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "のスタッフ情報を削除しました"})
}
