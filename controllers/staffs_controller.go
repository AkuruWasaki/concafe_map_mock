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
	staffs, err := models.Staffs(qm.Limit(5)).All(c, sc.DB)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, staffs)
}

// Create action: POST /staffs
func (sc StaffsController) Create(c *gin.Context) {
	var staff models.Staff

	if err := c.BindJSON(&staff); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := staff.Insert(c, sc.DB, boil.Infer()); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, staff)
}

// Update action: PUT /staffs/:id
func (sc StaffsController) Update(c *gin.Context) {
	// パスからID取得
	// idをint型に変換
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find staff data
	staff, err := models.FindStaff(c, sc.DB, idInt)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if staff == nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	// set param
	if err := c.BindJSON(&staff); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = staff.Update(c, sc.DB, boil.Infer())
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + strconv.Itoa(idInt) + "のスタッフ情報を更新しました"})
}

// Delete action: DELETE /staffs/:id
func (sc StaffsController) Delete(c *gin.Context) {
	// パスからID取得
	// idをint型に変換
	idInt, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// find staff data
	staff, err := models.FindStaff(c, sc.DB, idInt)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if staff == nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": "staff not found"})
		return
	}

	// delete staff
	if _, err = staff.Delete(c, sc.DB); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ID: " + strconv.Itoa(idInt) + "のスタッフ情報を削除しました"})
}
