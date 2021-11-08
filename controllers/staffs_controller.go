package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/akuruwasaki/concafe_map_mock/models"
	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

// Controller is staffs controller
type StaffsController struct {
	DB *sql.DB
}

// Index action: GET /staffs
func (sc StaffsController) Index(c *gin.Context) {
	db := sc.DB
	staffs, err := models.Staffs(qm.Limit(5)).All(c, db)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, staffs)
}

// Create action: POST /staffs
func (sc StaffsController) Create(c *gin.Context) {
	db := sc.DB
	var staff models.Staff

	if err := c.BindJSON(&staff); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := staff.Insert(c, db, boil.Infer()); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, staff)
}

// Update action: PUT /staffs/:id
func (sc StaffsController) Update(c *gin.Context) {
	db := sc.DB

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
	if err := c.BindJSON(&staff); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = staff.Update(c, db, boil.Infer())
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "のスタッフ情報を更新しました"})
}

// Delete action: DELETE /staffs/:id
func (sc StaffsController) Delete(c *gin.Context) {
	db := sc.DB

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
	if _, err = staff.Delete(c, db); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + id + "のスタッフ情報を削除しました"})
}
