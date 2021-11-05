package controller

import (
        "backend/entity"
        "github.com/gin-gonic/gin"
        "net/http"
)

// GET /Book
func ListBook(c *gin.Context) {
	var Book []entity.Book
	if err := entity.DB().Raw("SELECT * FROM Book").Scan(&Book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Book})
}


// GET /Book /id
func ListBook(c *gin.Context) {
	var Book entity.Book
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Book WHERE id = ?",id).Scan(&Book).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Book})
}

// POST /Book
func CreateBook(c *gin.Context) {
	var Book entity.Book
	if err := c.ShouldBindJSON(&Book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if err := entity.DB().Create(&Book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": Book})
}

// PATCH /Book
func UpdateBook(c *gin.Context) {
	var Book entity.Book
	if err := c.ShouldBindJSON(&Book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	if tx := entity.DB().Where("id = ?", Book.ID).First(&Book); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
			return
	}
	if err := entity.DB().Save(&Book).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Book})
}

// DELETE /Book/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM Book WHERE id = ?", id); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found"})
			return
	}
	c.JSON(http.StatusOK, gin.H{"data": Book})
}