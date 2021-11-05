package controller

import (
        "backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)

// GET /Admin
func ListAdmin(c *gin.Context) {
	var Admin []entity.Admin
	if err := entity.DB().Raw("SELECT * FROM Admin").Scan(&Admin).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Admin})
}


// GET /Admin /id
func ListAdmin(c *gin.Context) {
	var Admin entity.Admin
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Admin WHERE id = ?",id).Scan(&Admin).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Admin})
}

// POST /Admin
func CreateAdmin(c *gin.Context) {
	var Admin entity.Admin
	if err := c.ShouldBindJSON(&Admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&Admin).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Admin})
}

// PATCH /Admin
func UpdateAdmin(c *gin.Context) {
	var Admin entity.Admin
	if err := c.ShouldBindJSON(&Admin); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", Admin.ID).First(&Admin); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not found"})
			return
	}
	if err := entity.DB().Save(&Admin).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Admin})
}

// DELETE /Admin/:id
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Admin WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Admin not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}