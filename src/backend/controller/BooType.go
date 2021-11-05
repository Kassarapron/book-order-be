package controller

import (
        "backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)

// GET /BookType
func ListBookType(c *gin.Context) {
	var BookType []entity.BookType
	if err := entity.DB().Raw("SELECT * FROM BookType").Scan(&BookType).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookType})
}


// GET /BookType /id
func ListBookType(c *gin.Context) {
	var BookType entity.BookType
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM BookType WHERE id = ?",id).Scan(&BookType).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookType})
}

// POST /BookType
func CreateBookType(c *gin.Context) {
	var BookType entity.BookType
	if err := c.ShouldBindJSON(&BookType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&BookType).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookType})
}

// PATCH /BookType
func UpdateBookType(c *gin.Context) {
	var BookType entity.BookType
	if err := c.ShouldBindJSON(&BookType); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", BookType.ID).First(&BookType); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BookType not found"})
			return
	}
	if err := entity.DB().Save(&BookType).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookType})
}

// DELETE /BookType/:id
func DeleteBookType(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM BookType WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "BookType not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": BookType})
}